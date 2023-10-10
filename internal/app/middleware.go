package app

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// JWTMiddleware adalah middleware Gin untuk otentikasi JWT.
func JWTMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization header is missing",
			})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid authorization format",
			})
			c.Abort()
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Invalid signing method")
			}
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
			})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Failed to parse claims",
			})
			c.Abort()
			return
		}

		c.Set("claims", claims)
	}
}

// AdminMiddleware memeriksa apakah pengguna adalah admin berdasarkan klaim JWT.
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid or missing claims",
			})
			c.Abort()
			return
		}

		// Periksa apakah klaim "is_admin" ada dan bernilai true
		claimsMap, ok := claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Failed to parse claims",
			})
			c.Abort()
			return
		}

		isAdmin, ok := claimsMap["is_admin"].(bool)
		if !ok || !isAdmin {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Access denied. User is not an admin",
			})
			c.Abort()
			return
		}
	}
}
