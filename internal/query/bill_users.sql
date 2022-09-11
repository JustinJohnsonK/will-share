-- name: AddUserBill :one
INSERT INTO user_bills (bill_id, group_id, lend_user_id, borrow_user_id, amount)
VALUES ($1, $2, $3, $4, $5)
RETURNING bill_id, group_id, lend_user_id, borrow_user_id, amount;

-- name: SettleUserBillsByBillId :exec
UPDATE user_bills
SET is_active = False
WHERE bill_id = $1;

-- name: SettleUserBillsByGroupId :exec
UPDATE user_bills
SET is_active = False
WHERE group_id = $1;

-- name: SettleUserBillsByUserId :exec
UPDATE user_bills
SET is_active = False
WHERE (lend_user_id = $1
    AND borrow_user_id = $2)
    OR (lend_user_id = $2
    AND borrow_user_id = $1);