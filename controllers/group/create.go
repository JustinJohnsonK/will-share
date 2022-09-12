package group

import (
	"net/http"

	"github.com/JustinJohnsonK/will-share/internal/services"
	"github.com/JustinJohnsonK/will-share/internal/store"
	"github.com/JustinJohnsonK/will-share/internal/utils"
	"github.com/JustinJohnsonK/will-share/pkg/response"
	"github.com/labstack/echo/v4"
)

type createGroupRequest struct {
	GroupName   string `json:"group_name"`
	Description string `json:"description"`
}

func Create(s services.APIService) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var group createGroupRequest
		if err := c.Bind(&group); err != nil {
			return response.BadRequest(c)
		}

		newUser := store.CreateGroupParams{
			GroupName:   group.GroupName,
			Description: utils.ToNullString(group.Description),
		}

		i, err := s.GroupService.Create(ctx, newUser)
		if err != nil {
			return response.Unprocessable(c, nil)
		}

		return c.JSON(http.StatusCreated, i)
	}
}
