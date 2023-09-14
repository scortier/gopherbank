-- name: CreateAccount :one
INSERT INTO accounts (
  owner, 
  balance,
  currency
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetAccounts :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- many stands for multiple rows
-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1 --pagination, as data can be huge
OFFSET $2; -- no of records to skip

-- exec stands for execute
-- name: UpdateAccounts :exec
UPDATE accounts
SET balance = $2
WHERE id = $1;

-- name: DeleteAccounts :exec
DELETE FROM accounts
WHERE id = $1;
