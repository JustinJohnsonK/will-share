package app

import (
	"github.com/JustinJohnsonK/will-share/pkg/log"
)

var Logger log.Logger

func SetupLogger(service string, version string) {
	Logger = log.NewZerolog(service, version)
}
