package main

import (
	"os"

	"github.com/JustinJohnsonK/will-share/app"
	"github.com/JustinJohnsonK/will-share/internal/domain"
	"github.com/JustinJohnsonK/will-share/internal/services"
	"github.com/JustinJohnsonK/will-share/pkg/log"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	env := os.Getenv("WILLSHARE_ENV")

	if env == "" {
		env = "dev"
	}

	app.LoadConfig(env)
	app.SetupLogger("will-share", "0.1")

	deps := app.InitServiceDependencies()
	defer deps.Close()

	// e.Use(
	// 	middleware.Logger(app.Logger),

	// 	echoMiddleware.RequestID(),
	// 	echoMiddleware.RecoverWithConfig(echoMiddleware.RecoverConfig{
	// 		// TODO: Use app.Logger for structured logging
	// 		// or Create a recover middleware
	// 		LogLevel: echoLog.ERROR,
	// 	}),
	// 	echoMiddleware.CORS(),
	// )

	willshare := services.NewAPIService(deps)

	domain.Setup(*willshare, e)

	err := e.Start(app.Config.Server.Port)
	if err != nil {
		app.Logger.Panic("error while running server", err, log.Fields{})
	}
}
