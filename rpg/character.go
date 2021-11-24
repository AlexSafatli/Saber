package rpg

import "go.mongodb.org/mongo-driver/bson/primitive"

type Character struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Name            string
	Race            string
	Culture         string
	Profession      string
	Gender          string
	Sexuality       uint8
	Religion        string
	Characteristics map[string]EntityTag `json:"-"`
	Level           uint64
	GenerationDelta int    `json:"Generation Delta,omitempty"`
	YOBDelta        int    `json:"Year of Birth Delta,omitempty"`
	YODDelta        int    `json:"Year of Death Delta,omitempty"`
	DateDeltas      []Date `json:"Date Deltas,omitempty"`
	MotherID        uint   `json:"Mother ID,omitempty"`
	FatherID        uint   `json:"Father ID,omitempty"`
	LocationID      uint   `json:"Location ID,omitempty"`
	PC              bool   `json:"-"`
}

const (
	CharacterRaceDefault = "Human"

	CharacterSexualityHeterosexual = 0
	CharacterSexualityHomosexual   = 2
	CharacterSexualityBisexual     = 3
	CharacterSexualityAsexual      = 4
)

func NewCharacter(name, gender string) *Character {
	return &Character{Name: name, Gender: gender, Race: CharacterRaceDefault}
}
