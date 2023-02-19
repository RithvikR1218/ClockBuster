package main

import (
	//"errors"

	"clockbuster/src/api/controllers"
	"os"
	//"net/http"


	"github.com/gin-gonic/gin"
	"clockbuster/src/database/config"
)

func main() {

	r := gin.Default()
	config.InitDB()

	r.GET("/user", controllers.GetUsers) // new
	port := os.Getenv("PORT")
	r.Run("localhost:"+port)
}