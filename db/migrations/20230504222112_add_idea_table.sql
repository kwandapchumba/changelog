-- +goose Up
ALTER TABLE IF EXISTS idea DROP CONSTRAINT IF EXISTS fk_users;
ALTER TABLE IF EXISTS idea DROP CONSTRAINT IF EXISTS fk_company;
DROP TABLE IF EXISTS idea CASCADE;

CREATE TABLE idea (
    idea_id TEXT NOT NULL,
    idea_title TEXT NOT NULL,
    idea_description TEXT NOT NULL,
    idea_author TEXT NOT NULL,
    idea_created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    idea_updated_at TIMESTAMPTZ,
    idea_deleted_at TIMESTAMPTZ,
    company_id TEXT NOT NULL,
    PRIMARY KEY (idea_id),
    CONSTRAINT fk_users FOREIGN KEY (idea_author) REFERENCES "user" (user_id) ON DELETE CASCADE,
    CONSTRAINT fk_company FOREIGN KEY (company_id) REFERENCES company (company_id) ON DELETE CASCADE
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
ALTER TABLE IF EXISTS idea DROP CONSTRAINT IF EXISTS fk_users;
ALTER TABLE IF EXISTS idea DROP CONSTRAINT IF EXISTS fk_company;
DROP TABLE IF EXISTS idea CASCADE;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
