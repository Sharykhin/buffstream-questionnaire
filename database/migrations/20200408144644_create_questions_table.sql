-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE questions (
  id serial NOT NULL,
  uuid uuid NOT NULL,
  text text NOT NULL,
  created_at timestamp(0) NOT NULL,
  updated_at timestamp(0) NOT NULL,
  CONSTRAINT questions_pk PRIMARY KEY (uuid)
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE questions;