package routes

import (
	controller "restaurant_management_system/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	incomingRoutes.POST("/users/signup", controller.signup())
	incomingRoutes.POST("/users/signin", controller.signIn())
	// incomingRoutes.PUT("/users/:user_id", controllers.UpdateUser());
}
