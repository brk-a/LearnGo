package controllers

import (
	"context"
	"log"
	"net/http"
	"restaurant_management_system/database"
	"restaurant_management_system/helpers"
	"restaurant_management_system/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
        var orderItemPack OrderItemPack
        var order models.Order
        var orderItem models.OrderItem

        defer cancel()
        if err:=c.BindJSON(&orderItemPack); err!=nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        order.Order_date, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
        orderItemsToBeInserted := []interface{}{}
        order.Table_id = orderItemPack.Table_id
        order_id := OrderItemOrderCreator(order)
        for _, orderItem:=range orderItemPack.Order_items{
            orderItem.Order_id = order_id

            validationErr := validate.Struct(orderItem)
            defer cancel()
            if validationErr!=nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
                return
            }
        }
        var num = helpers.ToFixed(*orderItem.Unit_price, 2)
        orderItem.Unit_price = &num
        orderItem.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
        orderItem.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
        orderItem.ID = primitive.NewObjectID()
        orderItem.Order_item_id = orderItem.ID.Hex()
        orderItemsToBeInserted = append(orderItemsToBeInserted, orderItem)

        result, err := orderItemCollection.InsertMany(ctx, orderItemsToBeInserted)
        defer cancel()
        if err!= nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "error inserting order items"})
            return
        }

        defer cancel()
        c.JSON(http.StatusOK, result)
        
    }
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
			updateObj = append(updateObj, bson.E{"quantity", *&orderItem.Quantity})
		}
        if orderItem.Unit_price!= nil {
            updateObj = append(updateObj, bson.E{"unit_price", *&orderItem.Unit_price})
        }
        if orderItem.Food_id!= nil {
            updateObj = append(updateObj, bson.E{"food_id", *&orderItem.Food_id})
        }
        orderItem.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", orderItem.Updated_at})

        upsert := true
        opt := options.UpdateOptions{
            Upsert: &upsert,
        }
        result, err := orderItemCollection.UpdateOne(
            ctx,
            filter,
            bson.M{"$set": updateObj},
            &opt,
        )
        defer cancel()
        if err!= nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "error updating order item"})
            return
        }

        defer cancel()
        c.JSON(http.StatusOK, result)
	}
}
func DeleteOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
        orderItemId := c.Param("order_item_id")
        filter := bson.M{"orderItem_id": orderItemId}

        result, err := orderItemCollection.DeleteOne(ctx, filter)
        defer cancel()
        if err!= nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting order item"})
            return
        }
        if result.DeletedCount == 0 {
            c.JSON(http.StatusNotFound, gin.H{"error": "order item not found"})
            return
        }

        defer cancel()
        c.JSON(http.StatusOK, gin.H{"message": "order item deleted successfully"})
    }
}

func ItemsByOrder(id string) (OrderItems []primitive.M, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

    foodMatchStage := bson.D{
        {"$match", bson.D{
            {"order_id", id},
        }},
    }
    foodLookupStage := bson.D{
        {"$lookup", bson.D{
            {"from", "food"},
            {"localField", "food_id"},
            {"foreignField", "food_id"},
            {"as", "food"},
        }},
    }
    foodUnwindStage := bson.D{
        {"$unwind", bson.D{
            {"path", "$food"},
            {"preserveNullAndEmptyArrays", true},
        }},
    }
    orderLookupStage := bson.D{
        {"$lookup", bson.D{
            {"from", "order"},
            {"localField", "order_id"},
            {"foreignField", "order_id"},
            {"as", "order"},
        }},
    }
    orderUnwindStage := bson.D{
        {"$unwind", bson.D{
            {"path", "$order"},
            {"preserveNullAndEmptyArrays", true},
        }},
    }
    tableLookupStage := bson.D{
        {"$lookup", bson.D{
            {"from", "table"},
            {"localField", "order.table_id"},
            {"foreignField", "table_id"},
            {"as", "table"},
        }},
    }
    tableUnwindStage := bson.D{
        {"$unwind", bson.D{
            {"path", "$table"},
            {"preserveNullAndEmptyArrays", true},
        }},
    }
    projectStage1 := bson.D{
        {"$project", bson.D{
            {"id", 0},
            {"amount", "$food.price"},
            {"total_count", 1},
            {"food_name", "$food.name"},
            {"price", "$food.price"},
            {"food_image", "$food.image"},
            {"table_number", "$table.table_number"},
            {"table_id", "$table.table_id"},
            {"order_id", "$order.order_id"},
            {"order_date", "$order.order_date"},
            {"quantity", 1},
        }},
    }
    groupStage := bson.D{
        {"$group", bson.D{
            {"_id", bson.D{
                {"order_id", "$order_id"},
                {"table_id", "$table_id"},
                {"table_number", "$table_number"},
            }},
            {"payment_due", bson.D{
                {"$sum", "$amount"},
            }},
            {"total_count", bson.D{
                {"$sum", 1},
            }},
            {"order_items", bson.D{
                {"$push", "$$ROOT"},
            }},
        }},
    }
    projectStage2 := bson.D{
        {"$project", bson.D{
            {"id", 0},
            {"payment_due", 1},
            {"total_count", 1},
            {"table_number", "$_id.table_number"},
            {"order_items", 1},
        }},
    }

    result, err:= orderItemCollection.Aggregate(ctx, mongo.Pipeline{
        foodMatchStage,
        foodLookupStage,
        foodUnwindStage,
        orderLookupStage,
        orderUnwindStage,
        tableLookupStage,
        tableUnwindStage,
        projectStage1,
        groupStage,
        projectStage2,
    })
    defer cancel()
    if err!= nil {
        panic(err)
    }
    if err:=result.All(ctx, &OrderItems); err!=nil{
        panic(err)
    }

    defer cancel()
    return OrderItems, err
}
