// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: subscriptions.sql

package database

import (
	"context"
	"database/sql"
)

const createSubscription = `-- name: CreateSubscription :exec
INSERT INTO
  SUBSCRIPTIONS (
    account_id,
    service_name,
    plan_name,
    billing_frequency,
    cost,
    currency,
    icon,
    active,
    created_at,
    updated_at
  )
VALUES
  (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    1,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
  )
`

type CreateSubscriptionParams struct {
	AccountID        int32
	ServiceName      string
	PlanName         string
	BillingFrequency string
	Cost             float64
	Currency         string
	Icon             string
}

func (q *Queries) CreateSubscription(ctx context.Context, arg CreateSubscriptionParams) error {
	_, err := q.db.ExecContext(ctx, createSubscription,
		arg.AccountID,
		arg.ServiceName,
		arg.PlanName,
		arg.BillingFrequency,
		arg.Cost,
		arg.Currency,
		arg.Icon,
	)
	return err
}

const deleteSubscription = `-- name: DeleteSubscription :exec
DELETE FROM SUBSCRIPTIONS
WHERE
  id = ?
  AND account_id = ?
`

type DeleteSubscriptionParams struct {
	ID        int64
	AccountID int32
}

func (q *Queries) DeleteSubscription(ctx context.Context, arg DeleteSubscriptionParams) error {
	_, err := q.db.ExecContext(ctx, deleteSubscription, arg.ID, arg.AccountID)
	return err
}

const getSubscription = `-- name: GetSubscription :one
SELECT
  id,
  account_id,
  service_name,
  plan_name,
  billing_frequency,
  cost,
  currency,
  icon,
  active,
  created_at,
  updated_at
FROM
  SUBSCRIPTIONS
WHERE
  id = ?
`

func (q *Queries) GetSubscription(ctx context.Context, id int64) (Subscription, error) {
	row := q.db.QueryRowContext(ctx, getSubscription, id)
	var i Subscription
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.ServiceName,
		&i.PlanName,
		&i.BillingFrequency,
		&i.Cost,
		&i.Currency,
		&i.Icon,
		&i.Active,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listSubscriptions = `-- name: ListSubscriptions :many
SELECT
  id,
  account_id,
  service_name,
  plan_name,
  billing_frequency,
  cost,
  currency,
  icon,
  active,
  created_at,
  updated_at
FROM
  SUBSCRIPTIONS
WHERE
  account_id = ?
`

func (q *Queries) ListSubscriptions(ctx context.Context, accountID int32) ([]Subscription, error) {
	rows, err := q.db.QueryContext(ctx, listSubscriptions, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Subscription
	for rows.Next() {
		var i Subscription
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.ServiceName,
			&i.PlanName,
			&i.BillingFrequency,
			&i.Cost,
			&i.Currency,
			&i.Icon,
			&i.Active,
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

const updateSubscription = `-- name: UpdateSubscription :execresult
UPDATE SUBSCRIPTIONS
SET
  service_name = ?,
  plan_name = ?,
  billing_frequency = ?,
  cost = ?,
  currency = ?,
  icon = ?,
  active = ?,
  updated_at = CURRENT_TIMESTAMP
WHERE
  id = ?
  AND account_id = ?
`

type UpdateSubscriptionParams struct {
	ServiceName      string
	PlanName         string
	BillingFrequency string
	Cost             float64
	Currency         string
	Icon             string
	Active           sql.NullBool
	ID               int64
	AccountID        int32
}

func (q *Queries) UpdateSubscription(ctx context.Context, arg UpdateSubscriptionParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateSubscription,
		arg.ServiceName,
		arg.PlanName,
		arg.BillingFrequency,
		arg.Cost,
		arg.Currency,
		arg.Icon,
		arg.Active,
		arg.ID,
		arg.AccountID,
	)
}
