package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta"`
}

type ServerErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func Ok(ctx echo.Context, data interface{}, meta ...interface{}) error {
	return ctx.JSON(http.StatusOK, Response{Data: data, Success: true, Meta: meta})
}
