package data

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Models struct {
	LogEntry LogEntry
}

type LogEntry struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
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

func (l *LogEntry) Get(id string) (*LogEntry, error) {
	collection := client.Database("logs").Collection("logs")
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	var entry LogEntry
	err = collection.FindOne(context.TODO(), bson.M{"_id": ID}).Decode(&entry)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	return &entry, nil
}

func (l *LogEntry) Delete(id string) error {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	collection := client.Database("logs").Collection("logs")
	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": ID})
	if err != nil {
		return err
	}
	return nil
}

func (l *LogEntry) Update(entry LogEntry) (*mongo.UpdateResult, error) {
	ID, err := primitive.ObjectIDFromHex(entry.ID)
	if err != nil {
		return nil, err
	}
	collection := client.Database("logs").Collection("logs")
	updatedResult, err := collection.UpdateOne(context.TODO(), bson.M{"_id": ID}, bson.M{
		"$set": bson.M{
			"name":       l.Name,
			"data":       l.Data,
			"updated_at": time.Now(),
		},
	})
	if err != nil {
		return nil, err
	}
	return updatedResult, err
}

func (l *LogEntry) DropCollection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("logs").Collection("logs")

	if err := collection.Drop(ctx); err != nil {
		return err
	}

	return nil
}
