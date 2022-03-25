package main

import (
	"echoExample/controller"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strings"
)

func MyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fmt.Println(ctx.Request().Header.Get("Accept"))
		return next(ctx)
	}
}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookie, err := ctx.Cookie("userId")
		if err != nil {
			fmt.Println(err)
			if strings.Contains(err.Error(), "named cookie not present") {
				return ctx.String(http.StatusUnauthorized, "No cookies were sent")

			}
		}
		if cookie.Value == "user_id" {
			return next(ctx)
		}

		return ctx.String(http.StatusUnauthorized, "Wrong cookie")
	}
}

func main() {
	fmt.Println("hello world")

	ech := echo.New()
	ech.Use(MyMiddleware)
	ech.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "statusCode=${status}\n"}))

	ech.GET("/main", controller.GetMain)
	ech.GET("/user/:data", controller.GetUser)
	ech.POST("/user", controller.AddUser)

	adm := ech.Group("/admin")

	adm.GET("/main", controller.MainAdmin, checkCookie)
	adm.GET("/login", controller.LoginAdmin)

	ech.Start(":8080")
}
