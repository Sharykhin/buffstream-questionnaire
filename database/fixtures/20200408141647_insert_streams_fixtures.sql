-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
INSERT INTO streams(uuid, title, created_at, updated_at) VALUES
(uuid_generate_v4(), 'Netflix stream', now(), now()),
(uuid_generate_v4(), 'Youtube stream', now(), now()),
(uuid_generate_v4(), 'Scala stream', now(), now());
