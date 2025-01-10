// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: accounts.sql

package database

import (
	"context"
	"database/sql"
)

const createAccount = `-- name: CreateAccount :execresult
INSERT INTO ACCOUNTS (
    email,
    password_hash,
    first_name,
    last_name,
    profile_picture,
    created_at,
    updated_at
  )
VALUES (
    ?,
    ?,
    ?,
    ?,
    ?,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
  )
`

type CreateAccountParams struct {
	Email          string
	PasswordHash   string
	FirstName      string
	LastName       string
	ProfilePicture string
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createAccount,
		arg.Email,
		arg.PasswordHash,
		arg.FirstName,
		arg.LastName,
		arg.ProfilePicture,
	)
}

const deleteAccount = `-- name: DeleteAccount :exec
CALL delete_account (?)
`

func (q *Queries) DeleteAccount(ctx context.Context, AccountID int64) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, AccountID)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT id,
  email,
  first_name,
  last_name,
  profile_picture,
  created_at,
  updated_at
FROM ACCOUNTS
WHERE id = ?
`

type GetAccountRow struct {
	ID             int64
	Email          string
	FirstName      string
	LastName       string
	ProfilePicture string
	CreatedAt      sql.NullTime
	UpdatedAt      sql.NullTime
}

func (q *Queries) GetAccount(ctx context.Context, id int64) (GetAccountRow, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i GetAccountRow
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.ProfilePicture,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAccountByEmail = `-- name: GetAccountByEmail :one
SELECT id,
  email,
  first_name,
  last_name,
  password_hash,
  profile_picture,
  created_at,
  updated_at
FROM ACCOUNTS
WHERE email = ?
`

type GetAccountByEmailRow struct {
	ID             int64
	Email          string
	FirstName      string
	LastName       string
	PasswordHash   string
	ProfilePicture string
	CreatedAt      sql.NullTime
	UpdatedAt      sql.NullTime
}

func (q *Queries) GetAccountByEmail(ctx context.Context, email string) (GetAccountByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, getAccountByEmail, email)
	var i GetAccountByEmailRow
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.PasswordHash,
		&i.ProfilePicture,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT id,
  email,
  first_name,
  last_name,
  profile_picture,
  created_at,
  updated_at
FROM ACCOUNTS
ORDER BY created_at DESC
`

type ListAccountsRow struct {
	ID             int64
	Email          string
	FirstName      string
	LastName       string
	ProfilePicture string
	CreatedAt      sql.NullTime
	UpdatedAt      sql.NullTime
}

func (q *Queries) ListAccounts(ctx context.Context) ([]ListAccountsRow, error) {
	rows, err := q.db.QueryContext(ctx, listAccounts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListAccountsRow
	for rows.Next() {
		var i ListAccountsRow
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.FirstName,
			&i.LastName,
			&i.ProfilePicture,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const resetAutoIncrement = `-- name: ResetAutoIncrement :exec
ALTER TABLE ACCOUNTS AUTO_INCREMENT = 0
`

func (q *Queries) ResetAutoIncrement(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, resetAutoIncrement)
	return err
}

const updateAccount = `-- name: UpdateAccount :execresult
UPDATE ACCOUNTS
SET email = ?,
  password_hash = ?,
  first_name = ?,
  last_name = ?,
  profile_picture = ?,
  updated_at = CURRENT_TIMESTAMP
WHERE id = ?
`

type UpdateAccountParams struct {
	Email          string
	PasswordHash   string
	FirstName      string
	LastName       string
	ProfilePicture string
	ID             int64
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateAccount,
		arg.Email,
		arg.PasswordHash,
		arg.FirstName,
		arg.LastName,
		arg.ProfilePicture,
		arg.ID,
	)
}
