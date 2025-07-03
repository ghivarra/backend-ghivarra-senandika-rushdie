package router

import (
	"net/http"

	"github.com/ghivarra/app/module/controller/admin/product"
	"github.com/ghivarra/app/module/controller/auth"
	"github.com/ghivarra/app/module/controller/cart"
	"github.com/ghivarra/app/module/controller/order"
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
	adminProductRouterGroup.GET("/", product.Get)
	adminProductRouterGroup.POST("/create", roleCheckMiddleware.Run, product.Create)
	adminProductRouterGroup.PATCH("/update", roleCheckMiddleware.Run, product.Update)
	adminProductRouterGroup.DELETE("/delete", roleCheckMiddleware.Run, product.Delete)

	// cart group
	cartRouterGroup := router.Group("/cart")
	cartRouterGroup.Use(isLoggedOutMiddleware.Run)
	cartRouterGroup.GET("/", roleCheckMiddleware.Run, cart.Get)
	cartRouterGroup.POST("/add-product", roleCheckMiddleware.Run, cart.AddProduct)
	cartRouterGroup.POST("/buy", roleCheckMiddleware.Run, cart.Buy)

	// order group
	orderRouterGroup := router.Group("/order")
	orderRouterGroup.Use(isLoggedOutMiddleware.Run)
	orderRouterGroup.GET("/", roleCheckMiddleware.Run, order.Get)

	return router
}
