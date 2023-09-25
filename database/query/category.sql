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

-- name: UpdateCategory :one
UPDATE categorys
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categorys
WHERE id = $1;

