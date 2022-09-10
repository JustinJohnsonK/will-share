package services

import (
	"context"

	"github.com/JustinJohnsonK/will-share/internal/store"
)

type UserService struct {
	db *store.Queries
}

func NewUserService(db *store.Queries) *UserService {
	return &UserService{db: db}
}

func (s *UserService) Create(c context.Context, user store.CreateUserParams) (store.User, error) {
	i, err := s.db.CreateUser(c, user)
	return i, err
}

func (s *UserService) Get(c context.Context, user_id int64) (store.User, error) {
	i, err := s.db.GetUserById(c, user_id)
	return i, err
}
