package entities

type Character struct {
	Name            string
	Race            string
	Profession      string
	Characteristics map[string]EntityTag
	Level           uint64
}
