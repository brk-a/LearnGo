package routes

import (
	controller "restaurant_management_system/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/order_items", controller.GetOrderItems())
    incomingRoutes.GET("/order_items/:order_item_id", controller.GetOrderItem())
    incomingRoutes.POST("/order_items", controller.CreateOrderItem())
    incomingRoutes.PATCH("/order_items/:order_item_id", controller.UpdateOrderItem())
    // incomingRoutes.DELETE("/order_items/:order_item_id", controller.DeleteOrderItem())
    // incomingRoutes.GET("/order_items/order/:order_id", controller.GetOrderItemsByOrder())
    // incomingRoutes.GET("/order_items/food/:food_id", controller.GetOrderItemsByFood())
    // incomingRoutes.GET("/order_items/quantity/:quantity", controller.GetOrderItemsByQuantity())
}