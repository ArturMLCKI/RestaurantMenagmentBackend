package routes

import (
	controller "github.com/ArturMLCKI/RestaurantMenagmentBackend/tree/main/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:users_id", controller.GetUser())
	incomingRoutes.POST("/users/singup", controller.SignUp())
	incomingRoutes.POST("/users/login", controller.Login())
}
