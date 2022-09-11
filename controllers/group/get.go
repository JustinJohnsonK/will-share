package group

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/JustinJohnsonK/will-share/internal/services"
	"github.com/JustinJohnsonK/will-share/internal/store"
	"github.com/labstack/echo/v4"
)

type groupInfo struct {
	Group      store.GetGroupByGroupIdRow
	GroupUsers []store.GetGroupUsersByGroupIdRow
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
			fmt.Println("Group Details ", group_details, err)
			return err
		}

		group_users, err := s.GroupService.GetGroupUsers(ctx, group_id)
		if err != nil {
			fmt.Println("Group Users ", group_users, err)
			return err
		}

		group_info := groupInfo{
			Group:      group_details,
			GroupUsers: group_users,
		}

		return c.JSON(http.StatusOK, group_info)
	}
}
