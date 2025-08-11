package routes

import (
	"projects/go-rest-api/controllers"
	"projects/go-rest-api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	auth := router.Group("/auth")
	auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)

	authorized := router.Group("/auth")
	authorized.Use(middleware.AuthMiddleware())
	authorized.GET("/logout", controllers.Logout)
}
