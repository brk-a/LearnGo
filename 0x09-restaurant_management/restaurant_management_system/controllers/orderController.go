package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"restaurant_management_system/database"
	"restaurant_management_system/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")

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
            log.Fatal(err)
        }

        defer cancel()
        c.JSON(http.StatusOK, allOrders)
    }
}

func GetOrder() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
        orderId := c.Param("order_id")
        var order models.Order

        err := orderCollection.FindOne(ctx, bson.M{"order_id": orderId}).Decode(&order)
        defer cancel()
        if err!=nil {
            msg := fmt.Sprintf("error fetching order item")
            c.JSON(http.StatusNotFound, gin.H{"error": msg})
            return
        }

        defer cancel()
        c.JSON(http.StatusOK, order)
    }
}

func CreateOrder() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
        var order models.Order
        var table models.Table

        if err:=c.BindJSON(&order); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }

        validationError := validate.Struct(order)
        defer cancel()
        if validationError!=nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": validationError.Error()})
            return
        }

        if order.Table_id!=nil {
            err := tableCollection.FindOne(ctx, bson.M{"table_id": order.Table_id}).Decode(&table)
            defer cancel()
            if err!=nil {
                msg := fmt.Sprintf("error finding table that placed order")
                c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
                return
            }
        }

        order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
        order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
        order.ID = primitive.NewObjectID()
        order.Order_id = order.ID.Hex()

        result, err := orderCollection.InsertOne(ctx, order)
        defer cancel()
        if err!= nil {
            msg := fmt.Sprintf("error creating order")
            c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
            return
        }

        defer cancel()
        c.JSON(http.StatusOK, gin.H{"order": result})
    }
}

func UpdateOrder() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
        var order models.Order
        var table models.Table

        defer cancel()
        if err := c.BindJSON(&order); err!=nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        orderId := c.Param("order_id")
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
            bson.M{"$set": updateObj},
            &opt,
        )
        defer cancel()
        if err!= nil {
            msg := fmt.Sprintf("error updating order")
            c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
            return
        }

        defer cancel()
        c.JSON(http.StatusOK, result)
    }
}

func DeleteOrder() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
        orderId := c.Param("order_id")
        filter := bson.M{"order_id": orderId}

        result, err := orderCollection.DeleteOne(ctx, filter)
        defer cancel()
        if err!= nil {
            msg := fmt.Sprintf("error deleting order")
            c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
            return
        }

        if result.DeletedCount == 0 {
            c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
            return
        }

        defer cancel()
        c.JSON(http.StatusOK, gin.H{"message": "order deleted successfully"})
    }
}

func OrderItemOrderCreator(order models.Order) string {
    ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

    order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
    order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
    order.ID = primitive.NewObjectID()
    order.Order_id = order.ID.Hex()

    _, err := orderCollection.InsertOne(ctx, order)
    defer cancel()
    if err!= nil {
        msg := fmt.Sprintf("error creating order")
        return msg
    }
    
    defer cancel()
    return order.Order_id
}