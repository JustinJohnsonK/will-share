-- name: CreateGroup :one
INSERT INTO groups (group_name, description)
VALUES ($1, $2)
RETURNING group_id, group_name, description;

-- name: GetGroupByGroupId :one
SELECT group_id, group_name, description
FROM groups
WHERE group_id = $1
AND is_active = True;

-- name: DeleteGroupById :exec
UPDATE groups
SET is_active = False
WHERE group_id = $1;

-- name: HardDeleteUserById :exec
DELETE FROM groups 
WHERE group_id = $1;