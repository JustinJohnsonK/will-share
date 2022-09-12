-- name: AddBill :one
INSERT INTO bills (bill_title, bill_description, group_id, amount)
VALUES ($1, $2, $3, $4)
RETURNING bill_id, bill_title, bill_description, group_id, amount;

-- name: GetBillByBillId :one
SELECT bill_title, bill_description, group_id, amount
FROM bills
WHERE bill_id = $1;

-- name: DeleteBillByBillId :exec
UPDATE bills
SET is_active = False
WHERE bill_id = $1;

-- name: SettleBillByBillId :one
UPDATE bills
SET is_active = False
WHERE bill_id = $1
RETURNING amount;

-- name: SettleBillByGroupId :exec
UPDATE bills
SET is_active = False
WHERE group_id = $1;

-- name: DeleteBillsByGroupId :exec
UPDATE bills
SET is_active = False
WHERE group_id = $1;