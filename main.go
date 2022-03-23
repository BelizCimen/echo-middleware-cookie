package main

import (
	"echoExample/controller"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("hello world")

	ech := echo.New()
	ech.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "statusCode=${status}\n"}))

	ech.GET("/main", controller.GetMain)
	ech.GET("/user/:data", controller.GetUser)
	ech.POST("/user", controller.AddUser)

	adm := ech.Group("/admin")
	adm.GET("/main", controller.MainAdmin)

	ech.Start(":8080")
}
