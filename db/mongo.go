package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	mongoUri := "mongodb://user:112233@localhost:27017/todo?authSource=admin"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	db := client.Database("todoDB")

	err = createUniqueIndex(db.Collection("tasks"))
	if err != nil {
		return nil, err
	}

	return client.Database("todoDB"), nil
}
func createUniqueIndex(collection *mongo.Collection) error {
	ctx := context.Background()
	mod := mongo.IndexModel{
		Keys:    bson.D{{Key: "title", Value: 1}, {Key: "activeAt", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err := collection.Indexes().CreateOne(ctx, mod)
	return err
}
