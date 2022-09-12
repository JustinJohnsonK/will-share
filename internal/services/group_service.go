package services

import (
	"context"

	"github.com/JustinJohnsonK/will-share/internal/store"
)

type GroupService struct {
	db *store.Queries
}

func NewGroupService(db *store.Queries) *GroupService {
	return &GroupService{db: db}
}

func (s *GroupService) Create(c context.Context, group store.CreateGroupParams) (store.CreateGroupRow, error) {
	i, err := s.db.CreateGroup(c, group)
	return i, err
}

func (s *GroupService) Get(c context.Context, group_id int64) (store.GetGroupByGroupIdRow, error) {
	i, err := s.db.GetGroupByGroupId(c, group_id)
	return i, err
}

func (s *GroupService) Delete(c context.Context, group_id int64) error {
	err := s.db.DeleteGroupById(c, group_id)
	return err
}

func (s *GroupService) AddUserToGroup(c context.Context, userGroup store.AddUserToGroupParams) (store.AddUserToGroupRow, error) {
	i, err := s.db.AddUserToGroup(c, userGroup)
	return i, err
}

func (s *GroupService) GetGroupUsers(c context.Context, group_id int64) ([]store.GetGroupUsersByGroupIdRow, error) {
	i, err := s.db.GetGroupUsersByGroupId(c, group_id)
	return i, err
}

func (s *GroupService) GetGroupUserIds(c context.Context, group_id int64) ([]int64, error) {
	i, err := s.db.GetGroupUserIds(c, group_id)
	return i, err
}
