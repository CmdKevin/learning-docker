package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Get env
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Connect to database
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", dbUser, dbPass, dbName, dbHost, dbPort)
	fmt.Println("Connect to:", connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Connection to database failed:", err)
	}

	// Initialize Gin
	router := gin.Default()

	// Define a route for the "/" path
	router.GET("/", func(c *gin.Context) {
		// Print wise words
		c.String(http.StatusOK, "Wisdom is knowing what to do next; skill is knowing how to do it; and virtue is doing it. - David Starr Jordan")

		// Print to console that someone accessed the server
		fmt.Println("To Infinity and Beyond!")
	})

	// Define a route for "/ping" to ping the database
	router.GET("/ping", func(c *gin.Context) {
		// Test ping to database
		if err := db.Ping(); err != nil {
			c.String(http.StatusInternalServerError, "Failed to connect to the database: %v", err)
		} else {
			c.String(http.StatusOK, "Successfully connected to the database")
		}
	})

	// Start the server on port 78
	err = router.Run(":78")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
