-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE answers (
  id int NOT NULL,
  question_id int NOT NULL,
  text text NOT NULL,
  is_correct bool NOT NULL,
  created_at timestamp(0) NOT NULL,
  updated_at timestamp(0) NOT NULL,
  CONSTRAINT answers_pk PRIMARY KEY (question_id, id),
  CONSTRAINT answers_question_fk FOREIGN KEY (question_id) REFERENCES questions(id) ON DELETE RESTRICT ON UPDATE RESTRICT
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE answers;