-- name: GetUserById :one
SELECT user_id, user_name, phone_number
FROM users
WHERE user_id = $1
AND is_active = True;

-- name: CreateUser :one
INSERT INTO users (user_name, phone_number)
VALUES ($1, $2)
RETURNING user_id, user_name, phone_number;

-- name: DeleteUser :exec
UPDATE users
SET is_active = False
WHERE user_id = $1;

-- name: HardDeleteUser :exec
DELETE FROM users 
WHERE user_id = $1;