-- +goose Up
-- DROP CONSTRAINT IF EXISTS fk_user;
DROP TABLE IF EXISTS company CASCADE;

CREATE TABLE company (
    company_id TEXT UNIQUE NOT NULL,
    company_name TEXT NOT NULL,
    company_logo TEXT,
    user_id TEXT NOT NULL REFERENCES "user" (user_id) ON DELETE CASCADE,
    PRIMARY KEY (user_id)
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
-- DROP CONSTRAINT IF EXISTS fk_user;
DROP TABLE IF EXISTS company CASCADE;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
