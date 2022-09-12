package middleware

import (
	"time"

	"github.com/JustinJohnsonK/will-share/pkg/log"
	"github.com/labstack/echo/v4"
)

func Logger(logger log.Logger) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			request := c.Request()

			start := time.Now()
			err := next(c)

			if err != nil {
				c.Error(err)
			}

			stop := time.Now()

			latency := stop.Sub(start).String()
			response := c.Response()
			requestID := response.Header().Get(echo.HeaderXRequestID)

			fields := log.Fields{
				"request-id":   requestID,
				"status":       response.Status,
				"method":       request.Method,
				"path":         request.URL.Path,
				"latency":      latency,
				"query-params": request.URL.Query(),
			}

			// Do not log health check
			if c.Request().URL.Path != "/hea1thz" {
				logger.Info("info", fields)
			}

			return err
		}
	}
}
