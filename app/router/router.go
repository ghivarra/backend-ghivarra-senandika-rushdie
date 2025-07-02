package router

import (
	"net/http"

	"github.com/ghivarra/app/module/controller/admin/product"
	"github.com/ghivarra/app/module/controller/auth"
	userRole "github.com/ghivarra/app/module/controller/user-role"
	corsMiddleware "github.com/ghivarra/app/module/middleware/cors-middleware"
	isLoggedOutMiddleware "github.com/ghivarra/app/module/middleware/is-logged-out-middleware"
	roleCheckMiddleware "github.com/ghivarra/app/module/middleware/role-check-middleware"
	"github.com/gin-gonic/gin"
)

func RouteRegister() *gin.Engine {
	router := gin.Default()

	router.Use(corsMiddleware.Run())
	router.Use(gin.Recovery())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Layanan berfungsi secara normal.",
		})
	})

	// register & login
	authRouterGroup := router.Group("/auth")
	authRouterGroup.POST("/register", auth.Register)
	authRouterGroup.POST("/login", auth.Login)
	authRouterGroup.GET("/check", auth.Check)

	// role list
	router.GET("/user-role", userRole.Get)

	// admin group
	adminRouterGroup := router.Group("/admin")
	adminRouterGroup.Use(isLoggedOutMiddleware.Run)

	router.MaxMultipartMemory = 8 << 20
	adminProductRouterGroup := adminRouterGroup.Group("/product")
	adminProductRouterGroup.POST("create", roleCheckMiddleware.Run, product.Create)
	adminProductRouterGroup.PATCH("update", product.Create)
	adminProductRouterGroup.DELETE("delete", product.Create)

	return router
}
