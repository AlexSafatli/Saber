package entities

type Character struct {
	ID              uint
	Name            string
	Race            string
	Profession      string
	Gender          string
	Characteristics map[string]EntityTag `json:"-"`
	Level           uint64
	GenerationDelta int  `json:"Generation Delta"`
	YOBDelta        int  `json:"Year of Birth Delta"`
	YODDelta        int  `json:"Year of Death Delta"`
	MotherID        uint `json:"Mother ID"`
	FatherID        uint `json:"Father ID"`
	LocationID      uint `json:"Location ID"`
}

func NewCharacter(name, gender string) *Character {
	return &Character{Name: name, Gender: gender}
}
