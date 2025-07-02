package server

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func Run(router *gin.Engine) {
	serverPort := os.Getenv("PORT")
	serverHost := os.Getenv("HOST")

	router.Run(fmt.Sprintf("%s:%s", serverHost, serverPort))
}
