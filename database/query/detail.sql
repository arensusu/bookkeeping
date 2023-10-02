-- name: CreateDetail :one
INSERT INTO details (user_id, category_id, cost, date)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetDetailById :one
SELECT *
FROM details
WHERE id = $1;

-- name: ListDetailsByUserId :many
SELECT *
FROM details
WHERE user_id = $1
ORDER BY date ASC
LIMIT $2
OFFSET $3;

-- name: UpdateDetail :one
UPDATE details
SET category_id = $2, cost = $3, date = $4
WHERE id = $1
RETURNING *;

-- name: DeleteDetail :exec
DELETE FROM details
WHERE id = $1;
