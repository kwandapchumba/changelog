-- name: AddNewUser :one
INSERT INTO "user" (user_id, full_name, email, user_password)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: EmailExistsInDB :one
SELECT EXISTS (SELECT * FROM "user" WHERE email = $1);