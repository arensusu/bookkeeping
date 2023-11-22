-- name: CreateUser :one
INSERT INTO users (username, password, is_admin)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE username = $1
LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET password = $2
WHERE username = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE username = $1;
