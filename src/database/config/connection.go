package config

import (
)

var db *gorm.DB

func getDB () {
	return db
}

func Connect() {

	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
			panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Book{})
	if err != nil {
			return
	}
}