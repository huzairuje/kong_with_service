package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

var listUser = []interface{}{
	echo.Map{"id": 1, "name": "uje"},
	echo.Map{"id": 2, "name": "yazid"},
	echo.Map{"id": 3, "name": "ahmad"},
	echo.Map{"id": 4, "name": "dewo"},
	echo.Map{"id": 5, "name": "ipan"},
}

func selfRegisterKong() {
	name := "user-service"
	path := "http://localhost:1430"

	isSuccess, err := postKongAsService(name, path)
	if isSuccess {
		fmt.Print("succes register service to kong")
	} else {
		fmt.Print("error post Service", err)
	}

	isSuccessPostRoutesList, err := postKongAsRoute(name, "localhost", "GET", "/api/users")
	isSuccessPostRoutesDetail, err := postKongAsRoute(name, "localhost", "GET", "/api/users/1")
	if isSuccessPostRoutesList && isSuccessPostRoutesDetail {
		fmt.Print("succes register routes to kong ")
	} else {
		fmt.Print("error post Routes Body")
		fmt.Print("error post Routes", err)
	}
}

func main() {
	e := echo.New()
	api := e.Group("/api")
	userGroupRoute := api.Group("/users")
	userGroupRoute.GET("/1", func(context echo.Context) error {
		return context.JSON(http.StatusOK, echo.Map{
			"id": 1, "name": "uje"})
	})
	userGroupRoute.GET("", func(context echo.Context) error {
		return context.JSON(http.StatusOK, listUser)
	})

	go selfRegisterKong()

	e.Logger.Fatal(e.Start(":1430"))
}
