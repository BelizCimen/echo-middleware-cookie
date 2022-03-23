package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("hello world")

	ech := echo.New()

	ech.GET("/main", mainController)

	ech.Start(":8080")
}
