package main

import (
	controllers "cart/Controller"
	models "cart/Models"

	gin "github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/items", controllers.GetItems)
	router.POST("/items", controllers.CreateItem)

	router.Run()
}
