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

func GetStatus(s services.APIService) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		id := c.Param("id")

		fmt.Println("Group ID = ", id)

		group_id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			panic(err)
		}

		group_details, err := s.GroupService.Get(ctx, group_id)
		if err != nil {
			fmt.Println("Error Group Details ", group_details, err)
			// c.JSON(http.StatusNotFound, utils.Errors.MissingGroup)
			c.JSON(http.StatusNotFound, map[string]string{"error": "group not found"})
		}

		fmt.Println("Group Details ", group_details, err)

		// Borrowings by this user
		group_status, err := s.BillService.GetGroupStatusByGroupID(ctx, group_id)

		fmt.Printf("%v\n", group_status)

		return c.JSON(http.StatusOK, group_status)
	}
}
