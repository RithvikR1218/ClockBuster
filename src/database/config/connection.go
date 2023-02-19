package config

import (
	"clockbuster/src/database/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func getDB () *gorm.DB {
	return db
}

func InitDB() {
	godotenv.Load("/home/rithvik/spider/personal/clockbuster/src/env/.env")
	database, err := gorm.Open(mysql.Open("admin:password@(localhost:3306)/clockbuster?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
			panic("Failed to connect to database!")
	} else {
		println("Successfully connected to database")
	}
	err = database.AutoMigrate(&models.User{})
	if err != nil {
			return 
	}
	user := models.User {
		FirstName: "Rithvik",
		LastName: "Ravilla",
	}
	database.Create(&user)
}