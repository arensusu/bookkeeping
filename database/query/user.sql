-- name: CreateUser :one
INSERT INTO users (username, password, is_admin)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUserById :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET password = $2
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
