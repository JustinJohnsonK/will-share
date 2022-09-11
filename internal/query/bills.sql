-- name: AddBill :one
INSERT INTO bills (bill_title, bill_description, group_id, amount)
VALUES (bill_title = $1, 
    bill_description = $2,
    group_id = $3,
    amount = $4)
RETURNING *;

-- name: GetBillByBillId :one
SELECT *
FROM bills
WHERE bill_id = $1;

-- name: DeleteBillByBillId :exec
UPDATE bills
SET is_active = False
WHERE bill_id = $1;

-- name: DeleteBillsByGroupId :exec
UPDATE bills
SET is_active = False
WHERE group_id = $1;