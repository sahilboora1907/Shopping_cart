package main

import (
	controllers "cart/Controller"
	middleware "cart/Middlewares"
	models "cart/Models"

	gin "github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.POST("/users", controllers.Signup)
	router.POST("/users/login", controllers.Login)
	router.GET("/users", middleware.CheckAuth, controllers.GetUsers)
	router.GET("/items", controllers.GetItems)
	router.POST("/items", controllers.CreateItem)

	router.Run()
}
