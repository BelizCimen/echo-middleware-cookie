package controller

import (
	"echoExample/model"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

func GetUser(ctx echo.Context) error {
	dataType := ctx.Param("data")
	username := ctx.QueryParam("username")
	name := ctx.QueryParam("name")
	surname := ctx.QueryParam("surname")

	if dataType == "string" {
		return ctx.String(http.StatusOK, fmt.Sprintf("Username: %s, Name: %s, Surname: %s", username, name, surname))
	}

	if dataType == "json" {
		return ctx.JSON(http.StatusOK, map[string]string{
			"username": username,
			"name":     name,
			"surname":  surname,
		})
	}
	return ctx.String(http.StatusBadRequest, "json or string parameters only!")

}

func AddUser(ctx echo.Context) error {

	user := model.User{}

	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return err
	}

	fmt.Println(user)

	return ctx.String(200, "success")
}
