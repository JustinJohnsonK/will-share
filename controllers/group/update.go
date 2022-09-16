package group

import (
	"github.com/JustinJohnsonK/will-share/internal/services"
	"github.com/JustinJohnsonK/will-share/internal/store"
	"github.com/JustinJohnsonK/will-share/pkg/response"
	"github.com/labstack/echo/v4"
)

func AddUserToGroup(s services.APIService) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var userGroup store.AddUsersToGroupParams
		if err := c.Bind(&userGroup); err != nil {
			return response.BadRequest(c)
		}

		_, err := s.GroupService.AddUsersToGroup(ctx, userGroup)
		if err != nil {
			return response.BadRequest(c)
		}

		return response.Ok(c, nil)
	}
}
