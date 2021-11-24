package rpg

import "go.mongodb.org/mongo-driver/bson/primitive"

type Campaign struct {
	Name             string
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	StartDate        uint
	PlayerCharacters []*Character
}

func NewCampaign(name string) *Campaign {
	return &Campaign{Name: name}
}
