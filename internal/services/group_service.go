package services

import (
	"github.com/JustinJohnsonK/will-share/internal/store"
)

type GroupService struct {
	db *store.Queries
}

func NewGroupService(db *store.Queries) *GroupService {
	return &GroupService{db: db}
}
