package controllers

import (
	"net/http"

	//"clockbuster/src/database/config"
	"clockbuster/src/database/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetUsers (c *gin.Context) {
	var test models.User
	db.Find(&test)
	c.JSON(http.StatusOK, gin.H{"data": test})
}