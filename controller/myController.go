package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func SetHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fmt.Println(ctx.Request().Header.Get("Accept"))
		return next(ctx)
	}
}
