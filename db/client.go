package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

const (
	IndexDatabaseName            = "index"
	IndexCampaignsCollectionName = "campaigns"
	WorldsCollectionName         = "worlds"
)

func Connect() (*mongo.Client, context.Context, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		return nil, nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, nil, err
	}
	return client, ctx, nil
}

func Index(client *mongo.Client) *mongo.Database {
	return client.Database(IndexDatabaseName)
}

func Campaigns(client *mongo.Client) *mongo.Collection {
	return Index(client).Collection(IndexCampaignsCollectionName)
}

func Campaign(client *mongo.Client, id string) *mongo.Database {
	return client.Database(id)
}
