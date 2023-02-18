package main

import (
	//"errors"
	"fmt"
	"os"
	//"net/http"

	"github.com/gin-gonic/gin"
	"clockbuster/src/database/config"
)

func main() {

	r := gin.Default()
	//r.GET("/", func(c *gin.Context) {})
	port := os.Getenv("PORT")
	config.InitDB()
	fmt.Println("Server Runing on localhost:" + port)
	r.Run(":" + port)
	// fmt.Println("God damn it work bitch")
}