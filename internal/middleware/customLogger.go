package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func CustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		latency := time.Since(t)
		log.Printf("Method: %s | Status: %d | Latency: %s | Path: %s\n",
			c.Request.Method,
			c.Writer.Status(),
			latency,
			c.Request.URL.Path,
		)
	}
}
