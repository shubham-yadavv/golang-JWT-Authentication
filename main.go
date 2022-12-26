package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	routes "github.com/shubham-yadavv/golang-JWT-Authentication/routes"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	routes.AuthRoutes(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": "ok",
		})
	})

	router.Run(":" + port)

}
