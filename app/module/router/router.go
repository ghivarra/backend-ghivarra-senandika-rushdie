package router

import (
	"net/http"

	corsmiddleware "github.com/ghivarra/app/module/middleware/cors-middleware"
	"github.com/ghivarra/app/module/service"
	"github.com/gin-gonic/gin"
)

func RouteRegister() *gin.Engine {
	router := gin.Default()

	router.Use(corsmiddleware.Run())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Service is running normally.",
		})
	})

	router.GET("/product", service.ProductGet)
	router.GET("/user", service.UserGet)

	return router
}
