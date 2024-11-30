package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"restaurant_management_system/database"
	"restaurant_management_system/helpers"
	"restaurant_management_system/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")
var tableCollection *mongo.Collection = database.OpenCollection(database.Client, "table")

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		result, err := orderCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err!=nil {
            msg := fmt.Sprintf("error fetching orders")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

        var allOrders []bson.M
        if err = result.All(ctx, &allOrders); err!=nil {
            msg := fmt.Sprintf("error decoding orders")
            c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
            return
        }

        c.JSON(http.StatusOK, allOrders)
    }
}

func GetOrder() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
        orderId := c.Param(":order_id")
        var order models.Order

        err := orderCollection.FindOne(ctx, bson.M{"order_id": orderId}).Decode(&order)
        defer cancel()
        if err!=nil {
            msg := fmt.Sprintf("error fetching order item")
            c.JSON(http.StatusNotFound, gin.H{"error": msg})
            return
        }

        c.JSON(http.StatusOK, order)
    }
}

func CreateOrder() gin.HandlerFunc {
    return func(c *gin.Context) {}
}

func UpdateOrder() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
        var order models.Order
        var table models.Table

        if err := c.BindJSON(&order); err!=nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        orderId := c.Param(":order_id")
        filter := bson.M{"order_id": orderId}
        var updateObj primitive.D

        if order.Table_id!=nil {
            err := orderCollection.FindOne(ctx, bson.M{"table_id": order.Table_id}).Decode(&table)
            defer cancel()
            if err!=nil {
                msg := fmt.Sprintf("error finding table that placed order")
                c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
                return
            }
            updateObj = append(updateObj, bson.E{"table_id", order.Table_id})
        }
        order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", order.Updated_at})

        upsert := true
        opt := options.UpdateOptions{
            Upsert: &upsert,
        }
        result, err := orderCollection.UpdateOne(
            ctx,
            filter,
            bson.M{"$set": order},
            &opt,
        )
        if err!= nil {
            msg := fmt.Sprintf("error updating order")
            c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
            return
        }

        defer cancel()
        c.JSON(http.StatusOK, result)
    }
}