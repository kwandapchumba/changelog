-- +goose Up
DROP TABLE IF EXISTS changelog CASCADE;

CREATE TABLE changelog (
    id TEXT NOT NULL PRIMARY KEY,
    company_id TEXT NOT NULL REFERENCES company ON DELETE CASCADE,
    title VARCHAR NOT NULL,
    body TEXT NOT NULL,
    thumbnail TEXT NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMPTZ NOT NULL
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS changelog CASCADE;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
