-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE streams (
  id serial NOT NULL,
  uuid char(32) NOT NULL,
  title varchar(80) NOT NULL,
  created_at timestamp(0) NOT NULL,
  updated_at timestamp(0) NOT NULL,
  CONSTRAINT streams_pk PRIMARY KEY (uuid)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE streams;
