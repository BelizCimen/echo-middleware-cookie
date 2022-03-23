package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetMain(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Get request to main endpoint")
}
