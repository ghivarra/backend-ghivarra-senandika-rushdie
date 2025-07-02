package library

import (
	"math/rand"
	"os"
	"time"

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

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano())) // Seed the random number generator
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
