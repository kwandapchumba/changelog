// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: company.sql

package db

import (
	"context"
)

const addCompany = `-- name: AddCompany :one
INSERT INTO company (company_id, company_name, user_id)
VALUES ($1, $2, $3)
RETURNING company_id, company_name, company_logo, user_id
`

type AddCompanyParams struct {
	CompanyID   string `json:"company_id"`
	CompanyName string `json:"company_name"`
	UserID      string `json:"user_id"`
}

func (q *Queries) AddCompany(ctx context.Context, arg AddCompanyParams) (Company, error) {
	row := q.db.QueryRowContext(ctx, addCompany, arg.CompanyID, arg.CompanyName, arg.UserID)
	var i Company
	err := row.Scan(
		&i.CompanyID,
		&i.CompanyName,
		&i.CompanyLogo,
		&i.UserID,
	)
	return i, err
}