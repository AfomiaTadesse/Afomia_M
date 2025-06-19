package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// For now, just add a placeholder userID
		// In a real implementation, you would:
		// 1. Get token from header
		// 2. Validate JWT
		// 3. Extract userID
		// 4. Set in context
		c.Set("userID", "placeholder-user-id")
		c.Next()
	}
}