-- name: AddUserBill :one
INSERT INTO user_bills (bill_id, group_id, lend_user_id, borrow_user_id, amount)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;