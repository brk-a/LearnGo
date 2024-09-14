package controllers

import (
	"context"
	"log"
    "net/http"
	"restaurant_management_system/database"
	"restaurant_management_system/models"
    "time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
    "gopkg.in/mgo.v2/bson"
)

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

        result, err := menuCollection.Find(context.TODO(), bson.M{})
        defer  cancel()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching menus"})
            return
        }

        var allMenus []bson.M
        if err = result.All(ctx, &allMenus); err!=nil {
            log.Fatal(err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "error decoding results"})
            return
        }

        c.JSON(http.StatusOK, allMenus)
    };
}

func GetMenu() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
        menuId := c.Param("menu_id")
        var menu models.Menu

        err := menuCollection.FindOne(ctx, bson.M{"menuId": menuId}).Decode(&menu)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching menu"})
        }
        defer cancel()

        c.JSON(http.StatusOK, menu)
    };
}

func CreateMenu() gin.HandlerFunc {
    return func(c *gin.Context) {};
}

func UpdateMenu() gin.HandlerFunc {
    return func(c *gin.Context) {};
}