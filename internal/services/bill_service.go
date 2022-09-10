package services

import (
	"github.com/JustinJohnsonK/will-share/internal/store"
)

type BillService struct {
	db *store.Queries
}

func NewBillService(db *store.Queries) *BillService {
	return &BillService{db: db}
}
