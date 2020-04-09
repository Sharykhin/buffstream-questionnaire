package sql

import (
	"context"
	"database/sql"
	"fmt"

	"Sharykhin/buffstream-questionnaire/domains/question/repository/model"
	"Sharykhin/buffstream-questionnaire/errors"
)

type (
	// QuestionRepository is sql implementation that should satisfy repository interface
	QuestionRepository struct {
		db *sql.DB
	}
)

// FindOneByIDWithAnswers returns a specific question by its UUID and with all related answers
func (r *QuestionRepository) FindOneByIDWithAnswers(ctx context.Context, UUID string) (*model.Question, error) {
	query := `
	SELECT 
		q.id AS q_id, 
		q.uuid AS q_uuid,
		q.text AS q_text,
		q.created_at AS q_created_at,
		q.updated_at AS q_updated,
		a.id AS a_id,
		a.text AS a_text,
       	a.is_correct as a_is_correct,
		a.created_at AS a_created_at,
		a.updated_at AS a_updated_at
	FROM questions AS q 
	LEFT JOIN answers AS a ON a.question_id = q.id 
	WHERE q.uuid=$1`

	rows, err := r.db.QueryContext(ctx, query, UUID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query %v: %v", query, err)
	}

	var answers []model.Answer
	var question *model.Question

	for rows.Next() {
		var aggregate struct {
			model.Question
			model.Answer
		}
		err := rows.Scan(
			&aggregate.Question.ID,
			&aggregate.Question.UUID,
			&aggregate.Question.Text,
			&aggregate.Question.CreatedAt,
			&aggregate.Question.UpdatedAt,
			&aggregate.Answer.ID,
			&aggregate.Answer.Text,
			&aggregate.Answer.IsCorrect,
			&aggregate.Answer.CreatedAt,
			&aggregate.Answer.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan row into aggregated struct: %v", err)
		}

		if question == nil {
			question = &model.Question{
				ID:        aggregate.Question.ID,
				UUID:      aggregate.Question.UUID,
				Text:      aggregate.Question.Text,
				CreatedAt: aggregate.Question.CreatedAt,
				UpdatedAt: aggregate.Question.UpdatedAt,
				Answers:   answers,
			}
		}

		answer := model.Answer{
			ID:         aggregate.Answer.ID,
			QuestionID: aggregate.Question.ID,
			Text:       aggregate.Answer.Text,
			CreatedAt:  aggregate.Answer.CreatedAt,
			UpdatedAt:  aggregate.Answer.UpdatedAt,
		}

		answers = append(answers, answer)
	}

	if question == nil {
		return nil, fmt.Errorf("question was not found: %w", errors.ResourceNotFound)
	}

	question.Answers = answers

	return question, nil
}

// FindListByStreamID finds all questions associated with a specific stream
func (r *QuestionRepository) FindListByStreamID(ctx context.Context, UUID string) ([]model.Question, error) {
	query := `
	SELECT q.* FROM questions AS q
	INNER JOIN stream_questions AS sq ON  sq.question_id=q.id
	WHERE sq.stream_id=$1
	`
	rows, err := r.db.QueryContext(ctx, query, UUID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query %v: %v", query, err)
	}

	var questions []model.Question
	for rows.Next() {
		var question model.Question
		err := rows.Scan(&question.ID, &question.UUID, &question.Text, &question.CreatedAt, &question.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row into question struct: %v", err)
		}

		questions = append(questions, question)
	}

	return questions, rows.Err()

}

// NewQuestionRepository returns a new instance of sql question repository
// that should satisfy QuestionRepository interface
func NewQuestionRepository(db *sql.DB) *QuestionRepository {
	repo := QuestionRepository{
		db: db,
	}

	return &repo
}
