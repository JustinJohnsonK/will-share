package domain

import (
	"github.com/JustinJohnsonK/will-share/controllers/user"
	"github.com/JustinJohnsonK/will-share/internal/services"
	"github.com/JustinJohnsonK/will-share/internal/utils"
	"github.com/labstack/echo/v4"
)

func health(c echo.Context) error {
	return utils.Ok(c, nil)
}

func Setup(service services.APIService, e *echo.Echo) {

	e.POST("/user", user.Create(service))
	e.GET("/user/:id", user.Get(service))

	e.GET("/hea1thz", health)
}
