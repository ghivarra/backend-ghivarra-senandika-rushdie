package isloggedinmiddleware

import (
	"strings"

	"github.com/ghivarra/app/module/library/jwt"
	"github.com/gin-gonic/gin"
)

func Run(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	bearerToken := strings.Replace(authHeader, "Bearer ", "", 1)
	if bearerToken == "" {
		c.AbortWithStatusJSON(401, gin.H{
			"status":  "error",
			"message": "Anda harus login terlebih dahulu",
		})
		return
	}

	// validate
	valid, err := jwt.ValidateJWT(bearerToken)
	if err != nil || !valid {
		c.AbortWithStatusJSON(401, gin.H{
			"status":  "error",
			"message": "Anda harus login terlebih dahulu",
		})
		return
	}
}
