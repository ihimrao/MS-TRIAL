package data

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Models struct {
	LogEntry LogEntry
}

type LogEntry struct {
	Id        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string    `json:"name,omitempty" bson:"Name,omitempty"`
	Data      string    `json:"data,omitempty" bson:"data,omitempty"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

var client *mongo.Client

func New(mongo *mongo.Client) Models {
	client = mongo

	return Models{
		LogEntry: LogEntry{},
	}

}

func (l *LogEntry) Insert(entry LogEntry) error {
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), LogEntry{
		Name:      entry.Name,
		Data:      entry.Data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return err
	}
	return nil
}

func (l *LogEntry) Delete(id string) error {
	ID, err := primitive.ObjectIDFromHex(id)
	collection := client.Database("logs").Collection("logs")
	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": ID})
	if err != nil {
		return err
	}
	return nil
}

func (l *LogEntry) Edit(entry LogEntry) error {
	ID, err := primitive.ObjectIDFromHex(entry.Id)
	collection := client.Database("logs").Collection("logs")
	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": ID}, LogEntry{
		Name:      entry.Name,
		Data:      entry.Data,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return err
	}
	return nil
}
