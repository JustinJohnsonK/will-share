package group

import (
	"net/http"

	"github.com/JustinJohnsonK/will-share/internal/services"
	"github.com/JustinJohnsonK/will-share/internal/store"
	"github.com/labstack/echo/v4"
)

type addUserToGroupParams struct {
	GroupID int64   `json:"group_id"`
	UserID  []int64 `json:"user_id"`
}

func AddUserToGroup(s services.APIService) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var userGroup addUserToGroupParams
		if err := c.Bind(&userGroup); err != nil {
			return err
		}

		for _, id := range userGroup.UserID {
			user := store.AddUserToGroupParams{
				GroupID: userGroup.GroupID,
				UserID:  int64(id),
			}

			_, err := s.GroupService.AddUserToGroup(ctx, user)
			if err != nil {
				return err
			}
		}

		return c.JSON(http.StatusCreated, "")
	}
}
