package routes

import (
	"projects/go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func BookRoute(router *gin.Engine) {
	router.POST("/books", controllers.CreateBook)
	router.GET("/books", controllers.GetBooks)
	router.GET("/books/:id", controllers.GetBookByID)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)
}
