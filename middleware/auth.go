package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your_secret_key") // In production, use environment variables

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is required",
				"code":  "NO_AUTH_HEADER",
			})
			return
		}

		// If the token doesn't start with Bearer, try to add it
		if !strings.HasPrefix(strings.ToLower(authHeader), "bearer ") {
			authHeader = "Bearer " + authHeader
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization header format",
				"code":  "INVALID_AUTH_FORMAT",
			})
			return
		}

		tokenStr := bearerToken[1]
		if len(tokenStr) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token is empty",
				"code":  "EMPTY_TOKEN",
			})
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			// Validate the algorithm
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil {
			var errorMessage string
			var errorCode string

			switch {
			case err == jwt.ErrSignatureInvalid:
				errorMessage = "Invalid token signature"
				errorCode = "INVALID_SIGNATURE"
			case strings.Contains(err.Error(), "expired"):
				errorMessage = "Token has expired"
				errorCode = "TOKEN_EXPIRED"
			default:
				errorMessage = "Invalid token format or signature"
				errorCode = "INVALID_TOKEN"
			}

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": errorMessage,
				"code":  errorCode,
			})
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
				"code":  "INVALID_TOKEN",
			})
			return
		}

		// Additional expiration check
		if claims.ExpiresAt != nil {
			if claims.ExpiresAt.Before(time.Now()) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "Token has expired",
					"code":  "TOKEN_EXPIRED",
				})
				return
			}
		}

		// Store user information in context
		c.Set("username", claims.Subject)
		c.Next()
	}
}
