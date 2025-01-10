-- name: GetAccount :one
SELECT id,
  email,
  first_name,
  last_name,
  profile_picture,
  created_at,
  updated_at
FROM ACCOUNTS
WHERE id = ?;
-- name: GetAccountByEmail :one
SELECT id,
  email,
  first_name,
  last_name,
  password_hash,
  profile_picture,
  created_at,
  updated_at
FROM ACCOUNTS
WHERE email = ?;
-- name: ListAccounts :many
SELECT id,
  email,
  first_name,
  last_name,
  profile_picture,
  created_at,
  updated_at
FROM ACCOUNTS
ORDER BY created_at DESC;
-- name: CreateAccount :execresult
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
  );
-- name: UpdateAccount :execresult
UPDATE ACCOUNTS
SET email = ?,
  password_hash = ?,
  first_name = ?,
  last_name = ?,
  profile_picture = ?,
  updated_at = CURRENT_TIMESTAMP
WHERE id = ?;
-- name: DeleteAccount :exec
CALL delete_account (?);
-- name: ResetAutoIncrement :exec
ALTER TABLE ACCOUNTS AUTO_INCREMENT = 0;