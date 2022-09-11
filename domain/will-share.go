package domain

import (
	"github.com/JustinJohnsonK/will-share/controllers/group"
	"github.com/JustinJohnsonK/will-share/controllers/user"
	"github.com/JustinJohnsonK/will-share/internal/services"
	"github.com/JustinJohnsonK/will-share/internal/utils"
	"github.com/labstack/echo/v4"
)

func health(c echo.Context) error {
	return utils.Ok(c, nil)
}

func Setup(service services.APIService, e *echo.Echo) {

	// User endpoints
	e.POST("/user", user.Create(service))
	e.PUT("/user", user.Create(service))
	e.GET("/user/:id", user.Get(service))
	e.DELETE("/user/:id", user.Delete(service))

	// User endpoints
	e.POST("/group", group.Create(service))
	e.POST("/group/user", user.Create(service))
	e.GET("/group/:id", user.Get(service))
	e.DELETE("/group/:id", user.Delete(service))

	e.GET("/hea1thz", health)
}
