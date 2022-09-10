// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package store

import (
	"database/sql"
	"time"
)

type Bill struct {
	BillID          int64          `json:"bill_id"`
	GroupID         int64          `json:"group_id"`
	Amount          int32          `json:"amount"`
	BillTitle       sql.NullString `json:"bill_title"`
	BillDescription sql.NullString `json:"bill_description"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	IsActive        bool           `json:"is_active"`
}

type Group struct {
	GroupID     int64          `json:"group_id"`
	GroupName   string         `json:"group_name"`
	Description sql.NullString `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	IsActive    bool           `json:"is_active"`
}

type GroupUser struct {
	ID        int64     `json:"id"`
	GroupID   int64     `json:"group_id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsActive  bool      `json:"is_active"`
}

type User struct {
	UserID      int64          `json:"user_id"`
	UserName    string         `json:"user_name"`
	PhoneNumber sql.NullString `json:"phone_number"`
	IsActive    bool           `json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type UserBill struct {
	ID           int64     `json:"id"`
	LendUserID   int64     `json:"lend_user_id"`
	BorrowUserID int64     `json:"borrow_user_id"`
	BillID       int64     `json:"bill_id"`
	GroupID      int64     `json:"group_id"`
	Amount       int32     `json:"amount"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
