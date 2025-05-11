package main

import (
	"projects/go-rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.BookRoute(router)
	router.Run(":8080")
}
