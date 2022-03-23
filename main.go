package main

import (
	"echoExample/controller"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("hello world")

	ech := echo.New()

	ech.GET("/main", controller.GetMain)
	ech.GET("/user/:data", controller.GetUser)
	ech.POST("/user", controller.AddUser)

	ech.Start(":8080")
}
