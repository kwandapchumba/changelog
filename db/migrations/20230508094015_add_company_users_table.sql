-- +goose Up
DROP TABLE IF EXISTS company_user CASCADE;

CREATE TABLE company_user (
    user_id TEXT NOT NULL REFERENCES "user"(user_id),
    company_id TEXT NOT NULL REFERENCES company(company_id),
    PRIMARY KEY (user_id, company_id)
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS company_user CASCADE;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
