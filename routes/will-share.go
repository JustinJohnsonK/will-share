package routes

import (
	"github.com/JustinJohnsonK/will-share/controllers/bill"
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
	// Users
	e.POST("/user", user.Create(service))
	e.GET("/user/:id", user.Get(service))
	e.GET("/user/:id/status", user.GetStatus(service))
	e.DELETE("/user/:id", user.Delete(service))

	// Groups
	e.POST("/group", group.Create(service))
	e.GET("/group/:id", group.Get(service))
	e.DELETE("/group/:id", group.Delete(service))
	e.POST("/group/user", group.AddUserToGroup(service))
	e.GET("/group/:id/status", group.GetStatus(service))

	// Bills
	e.POST("/bill", bill.Create(service))
	e.PUT("/bill/settle", bill.SettleBill(service))
	e.PUT("/bill/group/settle", bill.SettleGroupBill(service))
	e.PUT("/bill/user/settle", bill.SettleUserBill(service))

	// Health Check
	e.GET("/hea1thz", health)
}
