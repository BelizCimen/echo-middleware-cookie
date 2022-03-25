package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func MainAdmin(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Admin endpoint")
}

func LoginAdmin(ctx echo.Context) error {
	username := ctx.QueryParam("username")
	password := ctx.QueryParam("password")
	if username == "admin" && password == "123" {
		cookie := http.Cookie{
			Name:    "userId",
			Value:   "user_id",
			Expires: time.Now().Add(48 * time.Hour),
		}
		ctx.SetCookie(&cookie)
		return ctx.String(http.StatusOK, "Logged in")
	}
	return ctx.String(http.StatusUnauthorized, "Wrong username or password")
}
