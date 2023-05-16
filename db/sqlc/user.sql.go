// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const addNewUser = `-- name: AddNewUser :one
INSERT INTO "user" (user_id, full_name, email, user_password)
VALUES ($1, $2, $3, $4)
RETURNING user_id, full_name, email, email_verified, picture, user_password, user_created_at, user_last_login
`

type AddNewUserParams struct {
	UserID       string         `json:"user_id"`
	FullName     sql.NullString `json:"full_name"`
	Email        string         `json:"email"`
	UserPassword string         `json:"user_password"`
}

func (q *Queries) AddNewUser(ctx context.Context, arg AddNewUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, addNewUser,
		arg.UserID,
		arg.FullName,
		arg.Email,
		arg.UserPassword,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.FullName,
		&i.Email,
		&i.EmailVerified,
		&i.Picture,
		&i.UserPassword,
		&i.UserCreatedAt,
		&i.UserLastLogin,
	)
	return i, err
}

const emailExistsInDB = `-- name: EmailExistsInDB :one
SELECT EXISTS (SELECT user_id, full_name, email, email_verified, picture, user_password, user_created_at, user_last_login FROM "user" WHERE email = $1)
`

func (q *Queries) EmailExistsInDB(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRowContext(ctx, emailExistsInDB, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}