-- name: AddCompany :one
INSERT INTO company (company_id, company_name, user_id)
VALUES ($1, $2, $3)
RETURNING *;