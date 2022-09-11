package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/JustinJohnsonK/will-share/internal/services"
	"github.com/labstack/echo/v4"
)

func Get(s services.APIService) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		userId := c.Get("user_id").(string)
		id := c.Param("id")

		fmt.Println("Get user by id", userId)
		user_id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			panic(err)
		}

		i, err := s.UserService.Get(ctx, user_id)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, i)
	}
}
