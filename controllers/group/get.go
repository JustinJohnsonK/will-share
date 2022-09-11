package group

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

		var id string
		id = c.Param("id")

		fmt.Println("Get group by id", id)
		group_id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			panic(err)
		}

		i, err := s.GroupService.Get(ctx, group_id)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, i)
	}
}
