package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func DebugTiming() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Debug") != "true" {
			c.Next()
			return
		}

		start := time.Now()
		c.Set("debugEnabled", true)
		c.Set("startTime", start)

		// Process request
		c.Next()

		end := time.Now()
		totalTime := end.Sub(start)

		// Store timing info to context for response shaping
		c.Set("requestTime", totalTime)
	}
}
