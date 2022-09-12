package group

import (
	"strconv"

	"github.com/JustinJohnsonK/will-share/internal/services"
	"github.com/JustinJohnsonK/will-share/internal/store"
	"github.com/JustinJohnsonK/will-share/pkg/response"
	"github.com/labstack/echo/v4"
)

type groupInfo struct {
	Group      store.GetGroupByGroupIdRow        `json:"group"`
	GroupUsers []store.GetGroupUsersByGroupIdRow `json:"group-users"`
}

func Get(s services.APIService) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var id string
		id = c.Param("id")

		group_id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			panic(err)
		}

		group_details, err := s.GroupService.Get(ctx, group_id)
		if err != nil {
			return response.NotFound(c)
		}

		group_users, err := s.GroupService.GetGroupUsers(ctx, group_id)
		if err != nil {
			return response.NotFound(c)
		}

		group_info := groupInfo{
			Group:      group_details,
			GroupUsers: group_users,
		}

		return response.Ok(c, group_info)
	}
}

func GetStatus(s services.APIService) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		id := c.Param("id")

		group_id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			panic(err)
		}

		// Validate group details
		_, err = s.GroupService.Get(ctx, group_id)
		if err != nil {
			return response.NotFound(c)
		}

		group_status, err := s.BillService.GetGroupStatusByGroupID(ctx, group_id)

		return response.Ok(c, group_status)
	}
}
