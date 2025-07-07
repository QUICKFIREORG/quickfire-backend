package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingHandler responds with "pong"
func PingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong from API handlers!")
}

// GetQuizHandler returns a dummy quiz
func GetQuizHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":        "quiz-123",
		"title":     "Go Programming Basics",
		"questions": []string{"What is a goroutine?", "What is a channel?"},
	})
}

// Add other handler functions here as needed
