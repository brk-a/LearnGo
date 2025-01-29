package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/brk-a/task_manager/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


// Database establishes a connection to the MongoDB database using the provided URI.
// It returns a pointer to the connected client and an error if any occurs during the connection process.
func Database() (client *mongo.Client, err error) {
	utils.LoadEnv()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)

	client, err = mongo.Connect(ctx, clientOptions)
	defer cancel()
	if err != nil {
		fmt.Println("error initialising DB", err)
		log.Panic(err)
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("error connecting to DB", err)
		log.Panic(err)
		return nil, err
	}

	fmt.Println("connected to DB...")
	return client, nil
}

var Client, err = Database()

// CloseDB closes the connection to the MongoDB database.
// It takes a pointer to the MongoDB client as a parameter and returns an error if any occurs during the disconnection process.
func CloseDB(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("disconnected from DB...")
	return nil
}

// GetCollection returns a reference to the specified collection in the connected MongoDB database.
// It takes a pointer to the MongoDB client and the name of the collection as parameters.
// The function also calls CloseDB to close the connection after retrieving the collection reference.
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection = client.Database("tasks").Collection(collectionName)
	// defer CloseDB(client)
	return collection
}
