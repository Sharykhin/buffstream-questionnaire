-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE streams (
  id serial NOT NULL,
  uuid uuid NOT NULL,
  title varchar(80) NOT NULL,
  created_at timestamp(0) NOT NULL,
  updated_at timestamp(0) NOT NULL,
  CONSTRAINT streams_pk PRIMARY KEY (id),
  CONSTRAINT streams_un UNIQUE (uuid)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE streams;
