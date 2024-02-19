package database

import (
	"context"
	"go-base-fs/utils"
	"log"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func DatabaseConnection() *mongo.Client {
	options := options.Client().ApplyURI(utils.GetEnvVar("MONGO_URI"))
	client, err := mongo.Connect(context.TODO(), options)

	if err != nil {
		log.Fatal("Error connecting to MongoDB", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Error connecting to MongoDB", err)
	}
	color.Cyan("MongoDB Connected Successfully")
	return client
}
