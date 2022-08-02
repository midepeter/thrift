// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: userquery.sql

package userstore

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  id,
  first_name,
  last_name,
  email,
  password,
  phone_number,
  created_at,
  updated_at
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, first_name, last_name, email, password, phone_number, created_at, updated_at
`

type CreateUserParams struct {
	ID          int32
	FirstName   string
	LastName    string
	Email       string
	Password    string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Password,
		arg.PhoneNumber,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.PhoneNumber,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getName = `-- name: GetName :one
SELECT id, first_name, last_name, email, password, phone_number, created_at, updated_at FROM users
WHERE id = $1 LIMIT $1
`

func (q *Queries) GetName(ctx context.Context, limit int32) (User, error) {
	row := q.db.QueryRow(ctx, getName, limit)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.PhoneNumber,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const removeUser = `-- name: RemoveUser :exec
DELETE FROM users WHERE id = $1
RETURNING id, first_name, last_name, email, password, phone_number, created_at, updated_at
`

func (q *Queries) RemoveUser(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, removeUser, id)
	return err
}