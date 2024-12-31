package controllers

import (
	"context"
	"log"
	"net/http"
	"restaurant_management_system/database"
	"restaurant_management_system/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderItemPack struct {
	Table_id    *string
	Order_items []models.OrderItem
}

var orderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "orderItem")

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		result, err := orderItemCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching order items"})
			return
		}

		var allOrderItems []bson.M
		if err = result.All(ctx, &allOrderItems); err != nil {
			log.Fatal(err)
		}

		defer cancel()
		c.JSON(http.StatusOK, allOrderItems)
	}
}

func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderId := c.Param("order_id")

		allOrderItems, err := ItemsByOrder(orderId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching order items by ID"})
			return
		}

		c.JSON(http.StatusOK, allOrderItems)
	}
}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		orderItemId := c.Param("order_item_id")
		var orderItem models.OrderItem

		err := orderItemCollection.FindOne(ctx, bson.M{"orderItem_id": orderItemId}).Decode(&orderItem)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "error decoding order item"})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, orderItem)
	}
}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var orderItem models.OrderItem

		defer cancel()
		if err := c.BindJSON(&orderItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		orderItemId := c.Param("order_item_id")
		filter := bson.M{"orderItem_id": orderItemId}
		var updateObj primitive.D
		if orderItem.Quantity != nil {
			updateObj = append(updateObj, bson.E{"quantity", orderItem.Quantity})
		}
	}
}
func DeleteOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func ItemsByOrder(id string) (OrderItems []primitive.M, err error) {
	return nil, nil
}
