-- +goose Up
-- DROP CONSTRAINT IF EXISTS fk_user;
DROP TABLE IF EXISTS company CASCADE;

CREATE TABLE company (
    company_id TEXT UNIQUE,
    company_name TEXT NOT NULL,
    company_belongs_to TEXT NOT NULL,
    company_logo TEXT,
    PRIMARY KEY (company_id, company_belongs_to),
    CONSTRAINT fk_user FOREIGN KEY (company_belongs_to) REFERENCES "user" (user_id) ON DELETE CASCADE
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
