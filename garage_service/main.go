package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func selfRegisterToKong() {
	name := "garage-service"
	path := "http://localhost:1431"

	isSuccess, err := postKongAsService(name, path)
	if isSuccess {
		fmt.Print("succes register service to kong")
	} else {
		fmt.Print("error post Service", err)
	}

	isSuccessPostRoutesList, err := postKongAsRoute(name, "localhost", "GET", "/api/garages")
	isSuccessPostRoutesDetail, err := postKongAsRoute(name, "localhost", "GET", "/api/garages/1")
	if isSuccessPostRoutesList && isSuccessPostRoutesDetail {
		fmt.Print("succes register routes to kong ")
	} else {
		fmt.Print("error post Routes", err)
	}
}

var listGarage = []interface{}{
	gin.H{"id": 1, "name": "garasi uje"},
	gin.H{"id": 2, "name": "garasi yazid"},
	gin.H{"id": 3, "name": "garasi ahmad"},
	gin.H{"id": 4, "name": "garasi dewo"},
	gin.H{"id": 5, "name": "garasi ipan"},
}

func main() {
	g := gin.Default()
	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "Not Found"})
	})
	api := g.Group("/api")
	garageGroupRoute := api.Group("/garages")
	garageGroupRoute.GET("/1", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"id": 1, "name": "garasi uje"})
		return
	})
	garageGroupRoute.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, listGarage)
		return
	})

	go selfRegisterToKong()

	_ = g.Run(":1431")

}
