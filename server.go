package main

import (
	"fmt"
	"os"
	//"net/http"
  	//"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	fmt.Println("Test")
	fmt.Println("God damn it work bitch")
}