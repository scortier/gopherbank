-- name: CreateEntry :one
INSERT INTO entries (
  account_id, 
  amount
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntry :many
SELECT * FROM entries
WHERE account_id = $1
ORDER BY id
LIMIT $2 
OFFSET $3;

-- -- exec stands for execute
-- -- name: UpdateEntry :exec
-- UPDATE entries
-- SET amount = $2
-- WHERE id = $1;

-- -- name: DeleteEntry :exec
-- DELETE FROM entries
-- WHERE id = $1;
