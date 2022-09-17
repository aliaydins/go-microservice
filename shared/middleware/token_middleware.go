package middleware

import (
	jwt_helper "github.com/aliaydins/microservice/shared/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("access_token") != "" {
			_, err := jwt_helper.VerifyToken(c.GetHeader("access_token"), secretKey)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
				c.Abort()
				return

			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized!"})
			c.Abort()
		}
		c.Next()
	}
}
