-- name: CreateCategory :one
INSERT INTO categorys (name)
VALUES ($1)
RETURNING *;

-- name: GetCategoryById :one
SELECT *
FROM categorys
WHERE id = $1
LIMIT 1;

-- name: ListCategorys :many
SELECT *
FROM categorys
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteCategory :exec
DELETE FROM categorys
WHERE id = $1;

