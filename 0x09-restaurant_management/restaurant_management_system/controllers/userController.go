package controllers

import (
	"restaurant_management_system/database"
	"restaurant_management_system/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var user models.User
var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO convert JSON data coming from DB/postman to sth Go understands
		// TODO validate data based on user struct
		// TODO check if email already exists
		// TODO check if phone number already exists
        // TODO generate verification code
        // TODO send verification code to user's email or phone number
        // TODO validate verification code from user's input
        // TODO update user's status to verified
        // TODO generate JWT token -> GenerateToken fn in helpers
        // TODO return JWT token
        // TODO hash password
		// TODO add metadata -> ID, updated_at, etc
        // TODO create new user in the DB
        // TODO send confirmation email
        // TODO return success message
	}
}

func SignIn() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func HashPassword(password string) string {
	return ""
}

func VerifyPassword(userPassword string, providePassword string) (bool, string) {
	return false, ""
}
