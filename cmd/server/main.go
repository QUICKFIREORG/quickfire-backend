package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/BerylCAtieno/quickfire-backend/internal/middleware"
	"github.com/BerylCAtieno/quickfire-backend/internal/routes"
)

func main() {
	// Set Gin to release mode for production, or debug mode for development.
	// gin.SetMode(gin.ReleaseMode) // Uncomment for production

	router := gin.Default()

	router.Use(middleware.CustomLogger())
	router.Use(gin.Recovery())

	apiV1 := router.Group("/api/v1")
	{

		routes.RegisterAPIRoutes(apiV1)
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Quickfire API!",
			"version": "1.0",
		})
	})

	// Start the server
	port := ":8080"
	fmt.Printf("Quickfire API server starting on %s\n", port)
	if err := router.Run(port); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
