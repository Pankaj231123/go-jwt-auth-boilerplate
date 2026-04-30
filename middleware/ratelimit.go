package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type visitor struct {
	count    int
	lastSeen time.Time
}

var (
	visitors = make(map[string]*visitor)
	mu       sync.Mutex
)

func RateLimit(maxRequests int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		mu.Lock()
		v, exists := visitors[ip]
		if !exists || time.Since(v.lastSeen) > window {
			visitors[ip] = &visitor{count: 1, lastSeen: time.Now()}
			mu.Unlock()
			c.Next()
			return
		}

		v.count++
		v.lastSeen = time.Now()

		if v.count > maxRequests {
			mu.Unlock()
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "too many requests, slow down",
			})
			c.Abort()
			return
		}

		mu.Unlock()
		c.Next()
	}
}