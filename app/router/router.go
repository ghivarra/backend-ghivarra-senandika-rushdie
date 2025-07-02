package router

import (
	"net/http"

	"github.com/ghivarra/app/module/controller/admin/product"
	"github.com/ghivarra/app/module/controller/auth"
	corsMiddleware "github.com/ghivarra/app/module/middleware/cors-middleware"
	isLoggedOutMiddleware "github.com/ghivarra/app/module/middleware/is-logged-out-middleware"
	"github.com/gin-gonic/gin"
)

func RouteRegister() *gin.Engine {
	router := gin.Default()

	router.Use(corsMiddleware.Run())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Service is running normally.",
		})
	})

	// register & login
	authRouterGroup := router.Group("/auth")
	authRouterGroup.POST("/register", auth.Register)
	authRouterGroup.POST("/login", auth.Login)
	authRouterGroup.GET("/check", auth.Check)

	// group and use middleware
	adminRouterGroup := router.Group("/admin")
	adminRouterGroup.Use(isLoggedOutMiddleware.Run)

	adminRouterGroup.POST("product/create", product.Create)

	return router
}
