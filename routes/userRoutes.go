package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shubham-yadavv/golang-JWT-Authentication/controllers"
	"github.com/shubham-yadavv/golang-JWT-Authentication/middleware"
)

func UserRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/users", middleware.Authenticate, middleware.AuthoriseRoles("admin"), controllers.GetAllUsers)
	incommingRoutes.GET("/user/:id", controllers.GetUserByID)
	incommingRoutes.GET("/me", middleware.Authenticate, controllers.GetProfile)
}
