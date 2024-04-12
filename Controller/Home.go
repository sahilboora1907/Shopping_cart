package controllers

import (
	http "net/http"

	gin "github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Home",
	})
}
