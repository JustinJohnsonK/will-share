-- name: AddUserToGroup :one
INSERT INTO group_users (group_id, user_id)
VALUES ($1, $2)
RETURNING *;

-- name: RemoveUserFromGroup :exec
UPDATE group_users
SET is_active = False
WHERE user_id = $1
AND group_id = $2;

-- name: HardRemoveUserFromGroup :exec
DELETE FROM group_users 
WHERE user_id = $1
AND group_id = $2;