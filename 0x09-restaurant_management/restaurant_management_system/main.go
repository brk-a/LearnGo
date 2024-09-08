package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"restaurant_management_system/database"
	"restaurant_management_system/routes"
	"restaurant_management_system/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food");

func main() {
	port := os.Getenv("PORT");
	if port == "" {
        port = "8080";
    }

	router := gin.New();
	router.Use(gin.Logger());

	routes.UserRoutes(router);
	router.Use(middleware.Authentication());

	routes.FoodRoutes(router);
	routes.OrderRoutes(router);
	routes.TableRoutes(router);
	routes.MenuRoutes(router);
	routes.OrderItemRoutes(router);
	routes.InvoiceRoutes(router);

	router.Run(":" + port);
}