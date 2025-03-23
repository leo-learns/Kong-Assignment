package main

import (
	"github.com/gin-gonic/gin"
	"services-api/services-api/db"
	"services-api/services-api/handlers"
)

func main() {
	// Initialize the database
	db.InitDB()

	// Set up Gin router
	r := gin.Default()

	// Define API endpoints
	r.GET("/services", handlers.GetServices)
	r.GET("/services/:id", handlers.GetService)
	r.GET("/services/:id/versions", handlers.GetServiceVersions)

	// Start the server on port 8080
	r.Run(":8080")
}
