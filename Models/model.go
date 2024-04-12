package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Username string `json:"username" unique:"true" binding:"required"`
	Password string `json:"password" binding:"required"`
	CartID   int    `json:"cart_id"`
}

type UsersInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Items struct {
	gorm.Model
	Name   string `json:"name"`
	Status string `json:"status"`
}

type ItemsInput struct {
	Name   string `json:"name" binding:"required"`
	Status string `json:"status" binding:"required"`
}

type Cart struct {
	gorm.Model
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
