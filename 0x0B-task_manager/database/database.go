package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Database() (client *mongo.Client, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)

	client, err = mongo.Connect(ctx, clientOptions)
	defer cancel()
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	
	fmt.Println("connected to mongoDB atlas...")
	return client, nil
}
