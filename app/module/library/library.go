package library

import (
	"os"

	"github.com/gin-gonic/gin"
)

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
