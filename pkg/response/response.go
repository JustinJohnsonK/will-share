package response

import (
	"net/http"

	"github.com/JustinJohnsonK/will-share/app"
	"github.com/JustinJohnsonK/will-share/pkg/log"
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

func Created(ctx echo.Context, data interface{}, meta ...interface{}) error {
	return ctx.JSON(http.StatusCreated, Response{Data: data, Success: true, Meta: meta})
}

func InternalError(ctx echo.Context, err error) error {
	app.Logger.Error("internal error", err, log.Fields{})
	return ctx.JSON(
		http.StatusInternalServerError,
		ServerErrorResponse{
			Success: false,
			Error:   "something went wrong. please try again",
		},
	)
}

func BadRequest(ctx echo.Context) error {
	app.Logger.Warn("bad request", log.Fields{})
	return ctx.JSON(
		http.StatusBadRequest,
		ServerErrorResponse{
			Success: false,
			Error:   "bad request",
		},
	)
}

func NotFound(ctx echo.Context) error {
	app.Logger.Warn("not found", log.Fields{})
	return ctx.JSON(
		http.StatusNotFound,
		ServerErrorResponse{
			Success: false,
			Error:   "not found",
		},
	)
}

func Unprocessable(ctx echo.Context, err error) error {
	return ctx.JSON(
		http.StatusUnprocessableEntity,
		ServerErrorResponse{
			Success: false,
			Error:   err.Error(),
		},
	)
}

func Forbidden(ctx echo.Context) error {
	app.Logger.Warn("forbidden", log.Fields{})
	return ctx.JSON(
		http.StatusForbidden,
		ServerErrorResponse{
			Success: false,
			Error:   "forbidden",
		},
	)
}
