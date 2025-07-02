package isloggedinmiddleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Run(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	requestToken := strings.Replace(authHeader, "Bearer ", "", 1)
	if requestToken == "" {
		c.AbortWithStatusJSON(401, gin.H{
			"status":  "error",
			"message": "Anda harus login terlebih dahulu",
		})
		return
	}

	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("failed to sign the JWT")
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(401, gin.H{
			"status":  "error",
			"message": "Anda harus login terlebih dahulu",
		})
		return
	}

	// check claim
	if claims, ok := token.Claims.(jwt.MapClaims); ok {

	} else {
		fmt.Println(err)
		c.AbortWithStatusJSON(401, gin.H{
			"status":  "error",
			"message": "Anda harus login terlebih dahulu",
		})
		return
	}
}
