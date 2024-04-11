package controllers

import (
	models "cart/Models"
	http "net/http"

	gin "github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context) {
	var items []models.Items
	models.DB.Find(&items)

	c.JSON(http.StatusOK, gin.H{"data": items})
}

func CreateItem(c *gin.Context) {
	var input models.ItemsInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item := models.Items{Name: input.Name, Status: input.Status}
	models.DB.Create(&item)

	c.JSON(http.StatusOK, gin.H{"data": item})
}
