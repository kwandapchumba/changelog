-- +goose Up
DROP TABLE IF EXISTS otp;

CREATE TABLE otp (
    otp TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    expiry TIMESTAMPTZ NOT NULL,
    verified BOOLEAN
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS otp;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
