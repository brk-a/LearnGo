package main

import (
	"log"

	"github.com/brk-a/task_manager/routes"
	"github.com/brk-a/task_manager/database"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)


var taskCollection *mongo.Collection
func main() {
	err := godotenv.Load(".env")
	if err!= nil {
        log.Fatal("Error loading.env file")
		return
    }

	client, dbErr := database.Database()
	if dbErr!=nil {
        log.Fatal("error initialising DB", err)
		return
    }

	taskCollection = client.Database("task_manager_db").Collection("tasks")
	routes.Routes()	
}