package services

import (
	"context"

	"github.com/JustinJohnsonK/will-share/internal/store"
)

type BillService struct {
	db *store.Queries
}

func NewBillService(db *store.Queries) *BillService {
	return &BillService{db: db}
}

func (s *BillService) Create(c context.Context, bill store.AddBillParams) (store.AddBillRow, error) {
	i, err := s.db.AddBill(c, bill)
	return i, err
}

func (s *BillService) CreateUserBill(c context.Context, bill store.AddUserBillParams) (store.AddUserBillRow, error) {
	i, err := s.db.AddUserBill(c, bill)
	return i, err
}

func (s *BillService) Get(c context.Context, bill_id int64) (store.GetBillByBillIdRow, error) {
	i, err := s.db.GetBillByBillId(c, bill_id)
	return i, err
}

func (s *BillService) GetBorrowingsByUserID(c context.Context, user_id int64) ([]store.GetBorrowingsByUserIdRow, error) {
	i, err := s.db.GetBorrowingsByUserId(c, user_id)
	return i, err
}

func (s *BillService) GetGroupStatusByGroupID(c context.Context, group_id int64) ([]store.GetGroupStatusByGroupIdRow, error) {
	i, err := s.db.GetGroupStatusByGroupId(c, group_id)
	return i, err
}

func (s *BillService) GetLendingsByUserID(c context.Context, user_id int64) ([]store.GetLendingsByUserIdRow, error) {
	i, err := s.db.GetLendingsByUserId(c, user_id)
	return i, err
}

func (s *BillService) DeleteBillByBillId(c context.Context, bill_id int64) error {
	err := s.db.DeleteBillByBillId(c, bill_id)
	return err
}

func (s *BillService) DeleteBillByGroupId(c context.Context, group_id int64) error {
	err := s.db.DeleteBillsByGroupId(c, group_id)
	return err
}

func (s *BillService) SettleBillByBillId(c context.Context, bill_id int64) (int32, error) {
	amount_settled, err := s.db.SettleBillByBillId(c, bill_id)
	if err != nil {
		return 0, err
	}
	err = s.db.SettleUserBillsByBillId(c, bill_id)
	return amount_settled, err
}

func (s *BillService) SettleBillByBillGroupId(c context.Context, group_id int64) error {
	err := s.db.SettleBillByGroupId(c, group_id)
	if err != nil {
		return err
	}
	err = s.db.SettleUserBillsByGroupId(c, group_id)
	return err
}

func (s *BillService) SettleBillByBillUserId(c context.Context, userIds store.SettleUserBillsByUserIdParams) error {
	err := s.db.SettleUserBillsByUserId(c, userIds)
	return err
}
