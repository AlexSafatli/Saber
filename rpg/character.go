package rpg

type Character struct {
	ID              uint `json:",omitempty"`
	Name            string
	Race            string
	Profession      string
	Gender          string
	Characteristics map[string]EntityTag `json:"-"`
	Level           uint64
	GenerationDelta int  `json:"Generation Delta,omitempty"`
	YOBDelta        int  `json:"Year of Birth Delta,omitempty"`
	YODDelta        int  `json:"Year of Death Delta,omitempty"`
	MotherID        uint `json:"Mother ID,omitempty"`
	FatherID        uint `json:"Father ID,omitempty"`
	LocationID      uint `json:"Location ID,omitempty"`
	PC              bool `json:"-"`
}

const (
	CharacterRaceDefault = "Human"
)

func NewCharacter(name, gender string) *Character {
	return &Character{Name: name, Gender: gender, Race: CharacterRaceDefault}
}
