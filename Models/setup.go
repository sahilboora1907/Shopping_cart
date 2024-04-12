package models

import (
	postgres "gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=Sahil#1907 dbname=Shopping_cart port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Items{}, &Cart{}, &Users{})
	DB = database
}
