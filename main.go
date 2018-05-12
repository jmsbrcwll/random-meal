package main

import (
	"github.com/gin-gonic/gin"
	"random-meal/app"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.GET("/meal", app.GetInstructions)
	}

	return router
}

func main() {
	router := SetupRouter()
	router.Run(":8080")
}