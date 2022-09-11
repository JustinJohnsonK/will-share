-- name: GetUserById :one
SELECT *
FROM users
WHERE user_id = $1;

-- name: CreateUser :one
INSERT INTO users (user_name, phone_number)
VALUES ($1, $2)
RETURNING *;

-- name: DeleteUser :exec
UPDATE users
SET is_active = False
WHERE user_id = $1;

-- name: HardDeleteUser :exec
DELETE FROM users WHERE user_id = $1;