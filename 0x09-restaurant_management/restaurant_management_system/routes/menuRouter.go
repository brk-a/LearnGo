package routes

import (
	controller "restaurant_management_system/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menu", controller.GetMenus())
	incomingRoutes.GET("/menu/:menu_id", controller.GetMenu())
	incomingRoutes.POST("/menu", controller.CreateMenu())
	incomingRoutes.PATCH("/menu/:menu_id", controller.UpdateMenu())
	// incomingRoutes.DELETE("/menu/:menu_id", controller.DeleteMenu())
	// incomingRoutes.GET("/menu/restaurant/:restaurant_id", controller.GetMenusByRestaurant())
	// incomingRoutes.GET("/menu/category/:category_id", controller.GetMenusByCategory())
	// incomingRoutes.GET("/menu/date/:start_date/:end_date", controller.GetMenusByDate())
	// incomingRoutes.GET("/menu/search/:search_query", controller.SearchMenus())
}
