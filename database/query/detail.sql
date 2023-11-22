-- name: CreateDetail :one
INSERT INTO details (username, category, cost, date)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetDetail :one
SELECT *
FROM details
WHERE id = $1;

-- name: ListDetailsByUser :many
SELECT *
FROM details
WHERE username = $1
ORDER BY date ASC
LIMIT $2
OFFSET $3;

-- name: UpdateDetail :one
UPDATE details
SET category = $2, cost = $3, date = $4
WHERE id = $1
RETURNING *;

-- name: DeleteDetail :exec
DELETE FROM details
WHERE id = $1;
