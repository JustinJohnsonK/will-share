-- name: CreateGroup :one
INSERT INTO groups (group_name, description)
VALUES ($1, $2)
RETURNING *;

-- name: GetGroupByGroupId :one
SELECT *
FROM groups
WHERE group_id = $1;

-- name: DeleteGroupById :exec
UPDATE groups
SET is_active = False
WHERE group_id = $1;

-- name: HardDeleteUserById :exec
DELETE FROM groups WHERE group_id = $1;