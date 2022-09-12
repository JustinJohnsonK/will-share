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

-- name: GetBorrowingsByUserId :many
SELECT lend_user_id, sum(amount) as amount
FROM user_bills
WHERE borrow_user_id = $1
AND is_active = True
GROUP BY lend_user_id;

-- name: GetLendingsByUserId :many
SELECT borrow_user_id, sum(amount) as amount
FROM user_bills
WHERE lend_user_id = $1
AND is_active = True
GROUP BY borrow_user_id;

-- name: GetGroupStatusByGroupId :many
SELECT user_bills.lend_user_id, u1.user_name as lend_user_name, user_bills.borrow_user_id, u2.user_name as borrow_user_name, sum(user_bills.amount) as amount
FROM user_bills
INNER JOIN users as u1 on u1.user_id = lend_user_id
INNER JOIN users as u2 on u2.user_id = borrow_user_id
WHERE user_bills.group_id = $1
AND user_bills.is_active = True
GROUP BY u1.user_name,
u2.user_name,
user_bills.lend_user_id,
user_bills.borrow_user_id;