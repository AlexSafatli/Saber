package entities

type Character struct {
	Name            string
	Profession      string
	Characteristics map[string]EntityTag
}
