package routes

import (
	"projects/go-rest-api/controllers"
	"projects/go-rest-api/middleware"

	"github.com/gin-gonic/gin"
)

func BookRoute(router *gin.Engine) {
	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	authorized.POST("/books", controllers.CreateBook)
	authorized.PUT("/books/:id", controllers.UpdateBook)
	authorized.DELETE("/books/:id", controllers.DeleteBook)

	// Public routes
	router.GET("/books", controllers.GetBooks)
	router.GET("/books/:id", controllers.GetBookByID)
}
