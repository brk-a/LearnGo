package controllers

import (
	"context"
	"log"
	"net/http"
	"restaurant_management_system/database"
	helpers "restaurant_management_system/helpers"
	"restaurant_management_system/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		_, err := userCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching users"})
			return
		}

		recordsPerPage, err := strconv.Atoi(c.Query("recordsPerpage"))
		if err != nil || recordsPerPage < 1 {
			recordsPerPage = 10
		}
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}
		startIndex := (page - 1) * recordsPerPage
		startIndex, err = strconv.Atoi(c.Query("start_index"))

		matchStage := bson.D{{
			"$match", bson.D{{}},
		},
		}
		projectStage := bson.D{{
			"$project", bson.D{
				{"_id", 0},
				{"total_count", 1},
				{"user_items", bson.D{
					{"$slice", []interface{}{"$data", startIndex, recordsPerPage}}},
				},
			},
		},
		}

		result, err := userCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, projectStage,
		})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching user items"})
			return
		}

		var allUsers []bson.M
		if err := result.All(ctx, &allUsers); err != nil {
			log.Fatal(err)
		}

		defer cancel()
		c.JSON(http.StatusOK, allUsers[0])
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		userId := c.Param("user_id")

		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error finding user"})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, user)
	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		// convert JSON data coming from DB/postman to sth Go understands
		defer cancel()
		if err:=c.BindJSON(&user); err!=nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// validate data based on user struct
		validationErr := validate.Struct(user)
		if validationErr!=nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
            return
        }
		// check if email already exists
		var count, err = userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err!=nil {
			log.Panic(err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while checking email"})
            return
        }
		// check if phone number already exists
		count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		defer cancel()
		if err!=nil {
			log.Panic(err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while checking phone number"})
            return
        }
		// email or phone number already exists
		if count>0{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "this email or phone number is already in use"})
			return
		}

		// hash password and add it to user object
		password := helpers.HashPassword(*user.Password)
		user.Password = &password
		// add metadata -> ID, updated_at, etc
		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()
		// generate JWT token -> GenerateToken fn in helpers
		token, refreshToken, _ := helpers.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, user.User_id)
		user.Token = &token
		user.Refresh_token = &refreshToken
		// create new user in the DB
		result, err := userCollection.InsertOne(ctx, user)
		defer cancel()
		if err!= nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating user"})
            return
        }

		// return success message
		defer cancel()
		c.JSON(http.StatusOK, result)

		// TODO generate verification code when email/phone number exists
		// TODO send verification code to user's email or phone number
		// TODO validate verification code from user's input
		// TODO update user's status to verified
		// TODO send confirmation email
	}
}

func SignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		var foundUser models.User
		// convert JSON data coming from DB/postman to sth Go understands
		if err:=c.BindJSON(&user); err!=nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
		// validate data based on user struct
		validationErr := validate.Struct(user)
		if validationErr!=nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
            return
        }
        // find user by email or phone number
        err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
        defer cancel()
        if err!= nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "email or phone number not found"})
            return
        }
		err = userCollection.FindOne(ctx, bson.M{"phone": user.Phone}).Decode(&foundUser)
        defer cancel()
        if err!= nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "email or phone number not found"})
            return
        }
        // verify password
        isMatch, msg := helpers.VerifyPassword(*foundUser.Password, *user.Password)
        if !isMatch {
            c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
            return
        }
		// generate token -> GenerateToken fn in helpers
		token, refreshToken, _ := helpers.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, foundUser.User_id)
		helpers.UpdateAllTokens(token, refreshToken, *&foundUser.User_id)

		// return success message
		defer cancel()
		c.JSON(http.StatusOK, foundUser)
	}
}


