package controllers

import (
	"clockbuster/src/database/config"
	"clockbuster/src/database/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = config.OpenCollection(config.Client,"user")

func GetUsers() gin.HandlerFunc {
	return func (c *gin.Context) {
		// var ctx, cancel = context.WithTimeout(context.Background(),100*time.Second)

	}
}

func GetUser() gin.HandlerFunc {
	return func (c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(),100*time.Second)
		userId := c.Param("user_id")
		var user models.User

		err := userCollection.FindOne(ctx,bson.M {"user_id": userId}).Decode(&user)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"error": "error while fetching users"})
		}
		c.JSON(http.StatusOK,user)
	}
}