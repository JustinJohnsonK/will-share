package main

import (
	"os"

	"github.com/JustinJohnsonK/will-share/app"
	"github.com/JustinJohnsonK/will-share/domain"
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
	app.SetupLogger("willshare", "0.1")

	deps := app.InitServiceDependencies()
	defer deps.Close()

	willshare := services.NewAPIService(deps)

	domain.Setup(*willshare, e)

	err := e.Start(app.Config.Server.Port)
	if err != nil {
		app.Logger.Panic("error while running server", err, log.Fields{})
	}
}
