-- name: GetSubscription :one
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
  id = ?;

-- name: ListSubscriptions :many
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
  account_id = ?;

-- name: CreateSubscription :exec
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
  );

-- name: UpdateSubscription :execresult
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
  AND account_id = ?;

-- name: DeleteSubscription :exec
DELETE FROM SUBSCRIPTIONS
WHERE
  id = ?
  AND account_id = ?;