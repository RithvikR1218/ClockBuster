package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"clockbuster/src/database/models"
)

var db *gorm.DB

func getDB () *gorm.DB {
	return db
}

func InitDB() {
	database, err := gorm.Open(mysql.Open("admin:password@(localhost:3306)/clockbuster?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
			panic("Failed to connect to database!")
	}
	err = database.AutoMigrate(&models.User{})
	if err != nil {
			return 
	}
}