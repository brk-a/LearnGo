package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"restaurant_management_system/database"
	"restaurant_management_system/helpers"
	"restaurant_management_system/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/bluesuncorp/validator.v5"
	"gopkg.in/mgo.v2/bson"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var validate = validator.New("validate", validator.BakedInValidators)

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var allFoods []bson.M
		recordsPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordsPerPage < 1 {
			recordsPerPage = 10
		}
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}
		skip := (page - 1) * recordsPerPage
		skip, err = strconv.Atoi(c.Query("skip"))
		matchStage := bson.D{{
			"$match", bson.D{{}},
		}}
		groupStage := bson.D{{
			"$group", bson.D{
				{"_id", bson.D{{"_id", "null"}}},
				{"total_count", bson.D{{"$sum", 1}}},
				{"$data", bson.D{{"$push", "$$ROOT"}}},
			},
		}}
		projectStage := bson.D{{
			"project", bson.D{
				{"_id", 0},
				{"total_count", 1},
				{"food_items", bson.D{{"$slice", []interface{}{"$data", skip, recordsPerPage}}}},
			},
		}}

		result, err := foodCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage,
			groupStage,
			projectStage,
		})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching food items"})
			return
		}
		if err = result.All(ctx, &allFoods); err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error decoding results"})
			return
		}

		c.JSON(http.StatusOK, allFoods[0])
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("food_id")
		var food models.Food

		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching food item"})
		}

		c.JSON(http.StatusOK, food)
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var food models.Food
		var menu models.Menu

		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if validationErr := validate.Struct(food); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)
		defer cancel()
		if err != nil {
			msg := fmt.Sprintf("menu was not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()
		var num = helpers.ToFixed(*food.Price, 2)
		food.Price = &num

		result, insertErr := foodCollection.InsertOne(ctx, food)
		if insertErr != nil {
			msg := fmt.Sprintf("error creating food item")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, result)
	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var food models.Food
		var menu models.Menu

		foodId := c.Param("food_id")
		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		menuId := c.Param("menu_id")
		filter := bson.M{"food_id": foodId}
		var updateObj primitive.D
		if food.Name != nil {
			updateObj = append(updateObj, bson.E{"name", food.Name})
		}
		if food.Price != nil {
			updateObj = append(updateObj, bson.E{"price", *food.Price})
		}
		if food.Food_image != nil {
			updateObj = append(updateObj, bson.E{"food_image", food.Food_image})
		}
		if food.Menu_id != nil {
			err := menuCollection.FindOne(ctx, bson.M{"menu_id": menuId}).Decode(&menu)
			defer cancel()
			if err!=nil{
				msg := fmt.Sprintf("error finding menu")
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}
			updateObj = append(updateObj, bson.E{"menu", *food.Menu_id})
		}

		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", food.Updated_at})

		upsert := true
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, err := foodCollection.UpdateOne(
			ctx,
			filter,
			bson.M{
				"$set": updateObj,
			},
			&opt,
		)
		if err != nil {
			msg := fmt.Sprintf("error updating food")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

// func round(num float64) int {
// 	return int(num)
// }

// func toFixed(num float64, precision int) float64 {
// 	return float64(num)
// }