package rpg

import "go.mongodb.org/mongo-driver/bson/primitive"

type EntityTag struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string
	Value string
}
