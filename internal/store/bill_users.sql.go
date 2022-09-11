// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: bill_users.sql

package store

import (
	"context"
)

const addUserBill = `-- name: AddUserBill :one
INSERT INTO user_bills (bill_id, group_id, lend_user_id, borrow_user_id, amount)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, lend_user_id, borrow_user_id, bill_id, group_id, amount, is_active, created_at, updated_at
`

type AddUserBillParams struct {
	BillID       int64 `json:"bill_id"`
	GroupID      int64 `json:"group_id"`
	LendUserID   int64 `json:"lend_user_id"`
	BorrowUserID int64 `json:"borrow_user_id"`
	Amount       int32 `json:"amount"`
}

func (q *Queries) AddUserBill(ctx context.Context, arg AddUserBillParams) (UserBill, error) {
	row := q.db.QueryRow(ctx, addUserBill,
		arg.BillID,
		arg.GroupID,
		arg.LendUserID,
		arg.BorrowUserID,
		arg.Amount,
	)
	var i UserBill
	err := row.Scan(
		&i.ID,
		&i.LendUserID,
		&i.BorrowUserID,
		&i.BillID,
		&i.GroupID,
		&i.Amount,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
