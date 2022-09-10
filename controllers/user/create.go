package user

import (
	"net/http"
	"strconv"

	"github.com/JustinJohnsonK/will-share/internal/services"
	"github.com/JustinJohnsonK/will-share/internal/store"
	"github.com/labstack/echo/v4"
)

func Create(s services.APIService) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		// Request body
		var user store.CreateUserParams
		if err := c.Bind(&user); err != nil {
			return err
		}

		// if err := s.SchoolService.ValidateCreate(ctx, school); err != nil {
		// 	return unprocessable(c, err)
		// }

		i, err := s.UserService.Create(ctx, user)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, i)
	}
}

func Get(s services.APIService) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var id string
		id = c.Param("id")

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
