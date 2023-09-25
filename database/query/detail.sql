-- name: CreateDetail :one
INSERT INTO details (user_id, category_id, cost, date)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: ListDetailsByUserId :many
SELECT *
FROM details
WHERE user_id = $1
ORDER BY date ASC
LIMIT $2
OFFSET $3;

-- name: DeleteDetail :exec
DELETE FROM details
WHERE id = $1;
