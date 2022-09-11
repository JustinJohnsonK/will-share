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

func (s *BillService) Create(c context.Context, group store.CreateGroupParams) (store.Group, error) {
	i, err := s.db.CreateGroup(c, group)
	return i, err
}

func (s *BillService) Get(c context.Context, group_id int64) (store.Bill, error) {
	i, err := s.db.GetBillByBillId(c, group_id)
	return i, err
}

func (s *BillService) Delete(c context.Context, bill_id int64) error {
	err := s.db.DeleteBillByBillId(c, bill_id)
	return err
}
