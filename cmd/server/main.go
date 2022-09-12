package main

import (
	"os"

	"github.com/JustinJohnsonK/will-share/app"
	"github.com/JustinJohnsonK/will-share/internal/middleware"
	"github.com/JustinJohnsonK/will-share/internal/services"
	"github.com/JustinJohnsonK/will-share/pkg/log"
	"github.com/JustinJohnsonK/will-share/routes"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	echoLog "github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	env := os.Getenv("WILLSHARE_ENV")

	if env == "" {
		env = "dev"
	}

	app.LoadConfig(env)

	app.SetupLogger("willshare", "0.1")

	deps := app.InitServiceDependencies()
	defer deps.Close()

	e.Use(
		middleware.Logger(app.Logger),

		echoMiddleware.RequestID(),
		echoMiddleware.RecoverWithConfig(echoMiddleware.RecoverConfig{
			LogLevel: echoLog.ERROR,
		}),
		echoMiddleware.CORS(),
	)

	willshare := services.NewAPIService(deps)

	routes.Setup(*willshare, e)

	err := e.Start(app.Config.Server.Port)

	if err != nil {
		app.Logger.Panic("error while running server", err, log.Fields{})
	}
}
