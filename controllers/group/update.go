package group

import (
	"fmt"
	"net/http"

	"github.com/JustinJohnsonK/will-share/internal/services"
	"github.com/JustinJohnsonK/will-share/internal/store"
	"github.com/labstack/echo/v4"
)

func AddUserToGroup(s services.APIService) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var userGroup store.AddUserToGroupParams
		if err := c.Bind(&userGroup); err != nil {
			return err
		}

		i, err := s.GroupService.AddUserToGroup(ctx, userGroup)
		if err != nil {
			fmt.Println("AddUserToGroup ", i, err)
			return err
		}

		return c.JSON(http.StatusCreated, i)
	}
}
