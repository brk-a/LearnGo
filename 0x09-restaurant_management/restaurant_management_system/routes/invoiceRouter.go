package routes

import (
	controller "restaurant_management_system/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", controller.GetInvoices())
    incomingRoutes.GET("/invoices/:invoice_id", controller.GetInvoice())
    incomingRoutes.POST("/invoices", controller.CreateInvoice())
    incomingRoutes.PATCH("/invoices/:invoice_id", controller.UpdateInvoice())
    // incomingRoutes.DELETE("/invoices/:invoice_id", controller.DeleteInvoice())
    // incomingRoutes.GET("/invoices/generate/:order_id", controller.GenerateInvoice())
    // incomingRoutes.GET("/invoices/total/:restaurant_id", controller.GetTotalInvoiceAmount())
    // incomingRoutes.GET("/invoices/restaurant/:restaurant_id", controller.GetInvoicesByRestaurant())
    // incomingRoutes.GET("/invoices/user/:user_id", controller.GetInvoicesByUser())
    // incomingRoutes.GET("/invoices/date/:start_date/:end_date", controller.GetInvoicesByDate())
}