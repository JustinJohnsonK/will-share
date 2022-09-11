package user

import (
	"strconv"

	"github.com/JustinJohnsonK/will-share/internal/services"
	"github.com/labstack/echo/v4"
)

func Delete(s services.APIService) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var id string
		id = c.Param("id")

		user_id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			panic(err)
		}

		err = s.UserService.Delete(ctx, user_id)
		if err != nil {
			return err
		}

		return nil
	}
}
