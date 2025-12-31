package middleware

import (
	"net/http"
	"strings"

	"github.com/cakeru/kanago-backend/internal/auth"
	"github.com/cakeru/kanago-backend/internal/config"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware verifies JWT tokens
func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get Authorization header
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Authorization header required",
            })
            c.Abort()
            return
        }

        // Parse Bearer token
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Invalid Authorization header format",
            })
            c.Abort()
            return
        }

        tokenString := parts[1]

        // Verify token
        claims, err := auth.VerifyToken(tokenString, cfg)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Invalid or expired token",
            })
            c.Abort()
            return
        }

        // Store user info in context
        c.Set("userID", claims.UserID)
        c.Set("email", claims.Email)
        c.Set("username", claims.Username)

        // Continue to next handler
        c.Next()
    }
}

// OptionalAuthMiddleware attempts to extract user info but doesn't require it
func OptionalAuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Authorization header
		authHeader := c.GetHeader("Authorization")

		//if no header, continue
		if authHeader == "" {
			c.Next()
			return
		}

		// Parse Bearer token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.Next()
			return
		}

		tokenString := parts[1]

		// try to verify token
		claims, err := auth.VerifyToken(tokenString, cfg)
		if err != nil {
			// Token invalid, continue without user info
			c.Next()
			return
		}

		// Token valid - store claims in context
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("username", claims.Username)

		// Continue to next handler
		c.Next()
	}
}