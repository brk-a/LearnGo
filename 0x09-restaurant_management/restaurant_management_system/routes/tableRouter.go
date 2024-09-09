package routes

import (
	controller "restaurant_management_system/controllers"

	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controller.GetTables())
	incomingRoutes.GET("/tables/:table_id", controller.GetTable())
	incomingRoutes.POST("/tables", controller.CreateTable())
	incomingRoutes.PATCH("/tables/:table_id", controller.UpdateTable())
	// incomingRoutes.DELETE("/tables/:table_id", controller.DeleteTable())
	// incomingRoutes.GET("/tables/restaurant/:restaurant_id", controller.GetTablesByRestaurant())
	// incomingRoutes.GET("/tables/status/:status", controller.GetTablesByStatus())
	// incomingRoutes.GET("/tables/capacity/:capacity", controller.GetTablesByCapacity())
	// incomingRoutes.GET("/tables/waiter/:waiter_id", controller.GetTablesByWaiter())
	// incomingRoutes.GET("/tables/customer/:customer_id", controller.GetTablesByCustomer())
}
