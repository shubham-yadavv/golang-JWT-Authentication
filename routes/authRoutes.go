package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shubham-yadavv/golang-JWT-Authentication/controllers"
	"github.com/shubham-yadavv/golang-JWT-Authentication/middleware"
)

func AuthRoutes(incommingRoutes *gin.Engine) {

	incommingRoutes.POST("/signup", controllers.Signup)
	incommingRoutes.POST("/login", controllers.Login)
	incommingRoutes.GET("/logout", controllers.Logout)
	incommingRoutes.GET("/user", middleware.Authenticate, middleware.AuthoriseRoles("user"), controllers.Validate)

}
