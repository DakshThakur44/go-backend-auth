package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LimitConcurrency(maxConcurrent int, maxWaitDuration time.Duration) gin.HandlerFunc {
	sem := make(chan struct{}, maxConcurrent)

	return func(c *gin.Context) {
		select {
		case sem <- struct{}{}:
			defer func() { <-sem }()
			c.Next()

		case <-time.After(maxWaitDuration):
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error": "Server is busy. Please try again later."})
			c.Abort()
		}
	}
}
