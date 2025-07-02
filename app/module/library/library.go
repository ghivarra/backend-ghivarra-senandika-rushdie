package library

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var JWTClaim jwt.MapClaims

func ErrorServer(message string, err error, c *gin.Context) {
	env := os.Getenv("ENV")

	if env == "development" {
		c.AbortWithError(500, err)
	} else {
		c.AbortWithStatusJSON(500, gin.H{
			"status":  "error",
			"message": message,
		})
	}
}
