package db

import (
	"context"
	"fmt"
	"github.com/AlexSafatli/Saber/rpg"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func CreateCampaign(client *mongo.Client, name string) (string, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	c := rpg.NewCampaign(name)
	index := Campaigns(client)
	res, err := index.InsertOne(ctx, c)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", res.InsertedID), nil
}
