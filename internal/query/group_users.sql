-- name: AddUserToGroup :one
INSERT INTO group_users (group_id, user_id)
VALUES ($1, $2)
RETURNING group_id, user_id;

-- name: GetGroupUsersByGroupId :many
SELECT group_users.user_id, users.user_name
FROM group_users
INNER JOIN users on users.user_id = group_users.user_id
WHERE group_users.group_id = $1
AND users.is_active = True;

-- name: RemoveUserFromGroup :exec
UPDATE group_users
SET is_active = False
WHERE user_id = $1
AND group_id = $2;

-- name: HardRemoveUserFromGroup :exec
DELETE FROM group_users 
WHERE user_id = $1
AND group_id = $2;

-- name: GetGroupUserIds :many
SELECT user_id
FROM group_users
WHERE group_id = $1
AND is_active = True;