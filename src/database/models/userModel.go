package models

type User struct {
	ID     		uint   `json:"id" gorm:"primary_key"`
	FirstName  	string `json:"FirstName"`
	LastName 	string `json:"LastName"`
  }