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
	"go.mongodb.org/mongo-driver/mongo/options"
)

var tableCollection *mongo.Collection = database.OpenCollection(database.Client, "table")

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

        result, err := tableCollection.Find(context.TODO(), bson.M{})
        defer cancel()
        if err!=nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching tables"})
            return
        }

        var allTables []bson.M
        if err = result.All(ctx, &allTables); err!=nil {
            log.Fatal(err)
        }

        defer cancel()
        c.JSON(http.StatusOK, allTables)
    }
}

func GetTable() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
        tableId := c.Param("table_id")
        var table models.Table

        err := tableCollection.FindOne(ctx, bson.M{"table_id": tableId}).Decode(&table)
        defer cancel()
        if err!=nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "error finding table"})
            return
        }

        defer cancel()
        c.JSON(http.StatusOK, table)
    }
}

func CreateTable() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
        var table models.Table

        defer cancel()
        if err := c.BindJSON(&table); err!=nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        defer cancel()
        valiationErr := validate.Struct(table)
        if valiationErr!= nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": valiationErr.Error()})
            return
        }

        table.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
        table.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
        table.ID = primitive.NewObjectID()
        table.Table_id = table.ID.Hex()

        result, err := tableCollection.InsertOne(ctx, table)
        if err!=nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "error inserting table"})
            return
        }

        defer cancel()
        c.JSON(http.StatusOK, gin.H{"table": result})
    }
}

func UpdateTable() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
        var table models.Table

        defer cancel()
        if err := c.BindJSON(&table); err!=nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        tableId := c.Param("table_id")
        var updateObj primitive.D
        filter := bson.M{"table_id": tableId}
        if table.Number_of_guests!=nil {
            updateObj = append(updateObj, bson.E{"number_of_guests", table.Number_of_guests})
        }
        if table.Table_number!=nil {
            updateObj = append(updateObj, bson.E{"table_number", table.Table_number})
        }
        table.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
        updateObj = append(updateObj, bson.E{"updated_at", table.Updated_at})

        upsert := true
        opt := options.UpdateOptions{
            Upsert: &upsert,
        }
        result, err := tableCollection.UpdateOne(
            ctx,
            filter,
            bson.M{"$set":  updateObj},
            &opt,
        )
        defer cancel()
        if err!=nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "error updating table"})
            return
        }

        defer cancel()
        c.JSON(http.StatusOK, result)
    }
}

func DeleteTable() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
        tableId := c.Param("table_id")

        filter := bson.M{"table_id": tableId}
        result, err := tableCollection.DeleteOne(ctx, filter)
        defer cancel()
        if err!= nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting table"})
            return
        }

        defer cancel()
        c.JSON(http.StatusOK, result)
    }
}