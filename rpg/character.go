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
	GenerationDelta int                `json:"Generation Delta,omitempty"`
	YOBDelta        int                `json:"Year of Birth Delta,omitempty"`
	YODDelta        int                `json:"Year of Death Delta,omitempty"`
	DateDeltas      []Date             `json:"Date Deltas,omitempty"`
	MotherID        primitive.ObjectID `json:"Mother ID,omitempty"`
	FatherID        primitive.ObjectID `json:"Father ID,omitempty"`
	LocationID      primitive.ObjectID `json:"Location ID,omitempty"`
	PC              bool               `json:"-"`
}

const (
	CharacterRaceDefault = "Human"

	CharacterSexualityHeterosexual = 0
	CharacterSexualityHomosexual   = 2
	CharacterSexualityBisexual     = 3
	CharacterSexualityAsexual      = 4
)

var (
	CharacterSexualities = []uint8{
		CharacterSexualityHeterosexual,
		CharacterSexualityHomosexual,
		CharacterSexualityBisexual,
		CharacterSexualityAsexual,
	}

	CharacterSexualitiesWeights = []uint{10, 4, 1, 1}
)

func NewCharacter(name, gender string) *Character {
	return &Character{Name: name, Gender: gender, Race: CharacterRaceDefault}
}
