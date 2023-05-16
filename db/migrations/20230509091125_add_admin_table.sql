-- +goose Up
DROP TABLE IF EXISTS "admin";

CREATE TABLE "admin"(
    user_id TEXT NOT NULL REFERENCES "user"(user_id) UNIQUE,
    company_id TEXT NOT NULL REFERENCES company(company_id),
    date_added TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS "admin";
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
