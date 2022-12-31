package main

import (
	"github.com/gin-gonic/gin"

	"github.com/shubham-yadavv/golang-JWT-Authentication/config"

	routes "github.com/shubham-yadavv/golang-JWT-Authentication/routes"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDB()
	config.SyncDatabase()
}

func main() {
	r := gin.Default()

	routes.AuthRoutes(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": "ok workkin",
		})
	})

	r.Run()

}
