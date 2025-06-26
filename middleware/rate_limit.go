package middleware

import (
	"net/http"
	"sync"

	"backend-server/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var visitors = make(map[string]*rate.Limiter)
var mu sync.Mutex

func getVisitor(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := visitors[ip]
	if !exists {
		rpsStr := utils.GetenvFloat("RATE_LIMIT_PER_SECOND", 1)
		burstStr := utils.GetenvInt("RATE_LIMIT_BURST", 5)
		limiter = rate.NewLimiter(rate.Limit(rpsStr), burstStr) // 1 request/sec, burst of 5
		visitors[ip] = limiter
	}
	return limiter
}

func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := getVisitor(ip)
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			return
		}
		c.Next()
	}
}
