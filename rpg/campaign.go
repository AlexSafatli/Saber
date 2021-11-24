package rpg

type Campaign struct {
	Name             string
	DM               string
	StartDate        uint
	PlayerCharacters []*Character
	World            *World
}
