package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/shubham-yadavv/golang-JWT-Authentication/controllers"
	"github.com/shubham-yadavv/golang-JWT-Authentication/middleware"
)

func UserRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.Use(middleware.Authenticate())
	incommingRoutes.GET("/users", controller.GetUsers())
	incommingRoutes.GET("users/:user_id", controller.GetUser())

}
