package main

import (
	"log"

	"github.com/brk-a/task_manager/routes"

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

	routes.Routes()	
}