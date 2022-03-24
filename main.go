package main

import (
	"echoExample/controller"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func MyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fmt.Println(ctx.Request().Header.Get("Accept"))
		return next(ctx)
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
	adm.Use(middleware.BasicAuth(func(username, password string, ctx echo.Context) (bool, error) {
		if username == "admin" && password == "123" {
			return true, nil
		}
		return false, nil
	}))

	adm.GET("/main", controller.MainAdmin)

	ech.Start(":8080")
}
