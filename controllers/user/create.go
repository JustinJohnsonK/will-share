package user

import (
	"github.com/JustinJohnsonK/will-share/internal/services"
	"github.com/JustinJohnsonK/will-share/internal/store"
	"github.com/JustinJohnsonK/will-share/internal/utils"
	"github.com/JustinJohnsonK/will-share/pkg/response"
	"github.com/labstack/echo/v4"
)

type createUserRequest struct {
	UserName    string `json:"user_name"`
	PhoneNumber string `json:"phone_number"`
}

func Create(s services.APIService) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var user createUserRequest
		if err := c.Bind(&user); err != nil {
			return response.BadRequest(c)
		}

		newUser := store.CreateUserParams{
			UserName:    user.UserName,
			PhoneNumber: utils.ToNullString(user.PhoneNumber),
		}

		i, err := s.UserService.Create(ctx, newUser)
		if err != nil {
			return response.BadRequest(c)
		}

		return response.Ok(c, i)
	}
}
