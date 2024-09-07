package routes

import (
	controller "restaurant_management_system/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controller.GetOrderItems())
    incomingRoutes.GET("/orderItems/:orderItem_id", controller.GetOrderItem())
    incomingRoutes.POST("/orderItems", controller.CreateOrderItem())
    incomingRoutes.PATCH("/orderItems/:orderItem_id", controller.UpdateOrderItem())
    // incomingRoutes.DELETE("/orderItems/:orderItem_id", controller.DeleteOrderItem())
    incomingRoutes.GET("/orderItems-order/:order_id", controller.GetOrderItemsByOrder())
    // incomingRoutes.GET("/orderItems-food/:food_id", controller.GetOrderItemsByFood())
    // incomingRoutes.GET("/orderItems-quantity/:quantity", controller.GetOrderItemsByQuantity())
}