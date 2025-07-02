package router

import (
	"net/http"

	"github.com/ghivarra/app/module/controller/auth"
	"github.com/ghivarra/app/module/controller/product"
	corsmiddleware "github.com/ghivarra/app/module/middleware/cors-middleware"
	isloggedinmiddleware "github.com/ghivarra/app/module/middleware/is-logged-in-middleware"
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
	router.GET("/user", service.UserGetAll)

	// register
	router.POST("/register", auth.Register)
	router.POST("/login", auth.Login)

	// group and use middleware
	adminRouterGroup := router.Group("/admin")
	adminRouterGroup.Use(isloggedinmiddleware.Run)

	adminRouterGroup.POST("product/create", product.Create)

	return router
}
