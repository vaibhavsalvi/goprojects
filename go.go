package main

import (
	"fmt"

	"helloworld/service"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Print("Program started...")
	router := gin.Default()
	// Define your routes and handlers here

	// Example route handler for the GET method
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	// Example route handler for the POST method
	router.POST("/submit", func(c *gin.Context) {
		// Handle the POST request
	})

	// Run the server on a specific address and port
	router.Run(":8080")
}

func getUser() service.User {
	userData := service.User{
		ID:       1,
		Username: "vaibhav",
		Email:    "v.g",
	}
	return userData
}
