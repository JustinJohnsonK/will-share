package services

import (
	"github.com/JustinJohnsonK/will-share/app"
	"github.com/JustinJohnsonK/will-share/internal/store"
)

type APIService struct {
	UserService  UserService
	GroupService GroupService
	BillService  BillService
}

func NewAPIService(deps *app.ServiceDependencies) *APIService {
	db := store.New(deps.Db)

	userService := *NewUserService(db)
	groupService := *NewGroupService(db)
	billService := *NewBillService(db)
	return &APIService{
		UserService:  userService,
		GroupService: groupService,
		BillService:  billService,
	}
}
