package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin
	router := gin.Default()

	// Define a route for the "/" path
	router.GET("/", func(c *gin.Context) {
		// Print wise words
		c.String(200, "Wisdom is knowing what to do next; skill is knowing how to do it; and virtue is doing it. - David Starr Jordan")

		// Print to console that someone accessed the server
		fmt.Println("Ouch!!")
	})

	// Start the server on port 78
	err := router.Run(":78")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
