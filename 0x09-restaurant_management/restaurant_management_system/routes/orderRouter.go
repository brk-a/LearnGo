package routes

import (
	controller "restaurant_management_system/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders", controller.GetOrders())
    incomingRoutes.GET("/orders/:order_id", controller.GetOrder())
    incomingRoutes.POST("/orders", controller.CreateOrder())
    incomingRoutes.PATCH("/orders/:order_id", controller.UpdateOrder())
    // incomingRoutes.DELETE("/orders/:order_id", controller.DeleteOrder())
    // incomingRoutes.GET("/orders/user/:user_id", controller.GetOrdersByUser())
    // incomingRoutes.GET("/orders/restaurant/:restaurant_id", controller.GetOrdersByRestaurant())
    // incomingRoutes.GET("/orders/date/:start_date/:end_date", controller.GetOrdersByDate())
}