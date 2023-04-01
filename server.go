package main

import (
	//"errors"
	"os"
	//"net/http"

	// "clockbuster/src/api/controllers"
	"clockbuster/src/api/routes"
	"clockbuster/src/database/config"
	// "clockbuster/src/middlewares"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
)

var UserTable *mongo.Collection = config.OpenCollection(config.Client,"User")

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	r := gin.Default()
	routes.UserRoutes(r)
	// routes.Use(middlewares.Authentication())
	
	r.Run("localhost:"+port)
}