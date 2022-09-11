// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: bills.sql

package store

import (
	"context"
	"database/sql"
)

const addBill = `-- name: AddBill :one
INSERT INTO bills (bill_title, bill_description, group_id, amount)
VALUES ($1, $2, $3, $4)
RETURNING bill_id, bill_title, bill_description, group_id, amount
`

type AddBillParams struct {
	BillTitle       sql.NullString `json:"bill_title"`
	BillDescription sql.NullString `json:"bill_description"`
	GroupID         int64          `json:"group_id"`
	Amount          int32          `json:"amount"`
}

type AddBillRow struct {
	BillID          int64          `json:"bill_id"`
	BillTitle       sql.NullString `json:"bill_title"`
	BillDescription sql.NullString `json:"bill_description"`
	GroupID         int64          `json:"group_id"`
	Amount          int32          `json:"amount"`
}

func (q *Queries) AddBill(ctx context.Context, arg AddBillParams) (AddBillRow, error) {
	row := q.db.QueryRow(ctx, addBill,
		arg.BillTitle,
		arg.BillDescription,
		arg.GroupID,
		arg.Amount,
	)
	var i AddBillRow
	err := row.Scan(
		&i.BillID,
		&i.BillTitle,
		&i.BillDescription,
		&i.GroupID,
		&i.Amount,
	)
	return i, err
}

const deleteBillByBillId = `-- name: DeleteBillByBillId :exec
UPDATE bills
SET is_active = False
WHERE bill_id = $1
`

func (q *Queries) DeleteBillByBillId(ctx context.Context, billID int64) error {
	_, err := q.db.Exec(ctx, deleteBillByBillId, billID)
	return err
}

const deleteBillsByGroupId = `-- name: DeleteBillsByGroupId :exec
UPDATE bills
SET is_active = False
WHERE group_id = $1
`

func (q *Queries) DeleteBillsByGroupId(ctx context.Context, groupID int64) error {
	_, err := q.db.Exec(ctx, deleteBillsByGroupId, groupID)
	return err
}

const getBillByBillId = `-- name: GetBillByBillId :one
SELECT bill_title, bill_description, group_id, amount
FROM bills
WHERE bill_id = $1
`

type GetBillByBillIdRow struct {
	BillTitle       sql.NullString `json:"bill_title"`
	BillDescription sql.NullString `json:"bill_description"`
	GroupID         int64          `json:"group_id"`
	Amount          int32          `json:"amount"`
}

func (q *Queries) GetBillByBillId(ctx context.Context, billID int64) (GetBillByBillIdRow, error) {
	row := q.db.QueryRow(ctx, getBillByBillId, billID)
	var i GetBillByBillIdRow
	err := row.Scan(
		&i.BillTitle,
		&i.BillDescription,
		&i.GroupID,
		&i.Amount,
	)
	return i, err
}

const settleBillByBillId = `-- name: SettleBillByBillId :one
UPDATE bills
SET is_active = False
WHERE bill_id = $1
RETURNING sum(amount)
`

func (q *Queries) SettleBillByBillId(ctx context.Context, billID int64) (int64, error) {
	row := q.db.QueryRow(ctx, settleBillByBillId, billID)
	var sum int64
	err := row.Scan(&sum)
	return sum, err
}

const settleBillByGroupId = `-- name: SettleBillByGroupId :one
UPDATE bills
SET is_active = False
WHERE group_id = $1
RETURNING sum(amount)
`

func (q *Queries) SettleBillByGroupId(ctx context.Context, groupID int64) (int64, error) {
	row := q.db.QueryRow(ctx, settleBillByGroupId, groupID)
	var sum int64
	err := row.Scan(&sum)
	return sum, err
}
