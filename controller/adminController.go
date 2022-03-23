package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func MainAdmin(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Admin endpoint")
}
