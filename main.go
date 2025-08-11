package main

import (
	"log"
	"projects/go-rest-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	router := gin.Default()
	routes.BookRoute(router)
	routes.UserRoute(router)
	router.Run(":8080")
}
