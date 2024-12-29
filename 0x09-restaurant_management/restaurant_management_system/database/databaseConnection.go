package database

import (
    "context"
    "fmt"
    "log"
	"time"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	MongoDb := "mongodb://localhost:" + os.Getenv("MONGODB_PORT"); //27107
	fmt.Println(MongoDb);

	// TODO: mongo.NewClient is deprecated: Use [Connect] instead
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb));
	if err != nil {
		log.Fatal(err);
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second);
	defer cancel();

	err = client.Connect(ctx);
	if err != nil {
		log.Fatal(err);
	}

	fmt.Println("connected to DB");

	return client;
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant").Collection(collectionName);

	return collection;
}