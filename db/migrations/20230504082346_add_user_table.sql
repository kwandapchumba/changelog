-- +goose Up
DROP TABLE IF EXISTS "user" CASCADE;

CREATE TABLE "user" (
    user_id TEXT UNIQUE,
    full_name TEXT,
    email TEXT NOT NULL CONSTRAINT email_must_be_unique UNIQUE,
    email_verified BOOLEAN NOT NULL DEFAULT 'false',
    picture TEXT NOT NULL DEFAULT '',
    user_password TEXT NOT NULL,
    user_created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_last_login TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, email)
);

CREATE INDEX email_idx ON "user"(LOWER(email));
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS "user" CASCADE;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
