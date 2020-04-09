-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE stream_questions (
   stream_id int NOT NULL,
   question_id int NOT NULL,
   created_at timestamp(0) NOT NULL,
   CONSTRAINT stream_questions_pk PRIMARY KEY (stream_id, question_id),
   CONSTRAINT stream_questions_stream_fk FOREIGN KEY (stream_id) REFERENCES streams(id) ON DELETE RESTRICT ON UPDATE RESTRICT,
   CONSTRAINT stream_questions_question_fk FOREIGN KEY (question_id) REFERENCES questions(id) ON DELETE RESTRICT ON UPDATE RESTRICT
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE stream_questions;
