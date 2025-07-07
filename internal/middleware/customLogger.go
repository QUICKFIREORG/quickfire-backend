package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// CustomLogger is a simple middleware to log request details.
func CustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Process request
		c.Next()

		// After request
		latency := time.Since(t)
		log.Printf("Method: %s | Status: %d | Latency: %s | Path: %s\n",
			c.Request.Method,
			c.Writer.Status(),
			latency,
			c.Request.URL.Path,
		)
	}
}

// Add other middleware functions here as needed
