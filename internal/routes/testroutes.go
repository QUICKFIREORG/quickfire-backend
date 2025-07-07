package routes

import (
	"github.com/BerylCAtieno/quickfire-backend/internal/handlers" // Import your handlers package

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes sets up the API routes for a given Gin router group.
func RegisterAPIRoutes(rg *gin.RouterGroup) {
	// Example API routes
	rg.GET("/ping", handlers.PingHandler)
	rg.GET("/quiz", handlers.GetQuizHandler)

	// Add more API routes here, e.g.,
	// rg.POST("/users", handlers.CreateUserHandler)
	// rg.GET("/users/:id", handlers.GetUserHandler)
}

// You can define other route registration functions if you have different route sets
// func RegisterAdminRoutes(rg *gin.RouterGroup) { ... }
