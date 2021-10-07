package rpg

type Campaign struct {
	Name             string
	DM               string
	PlayerCharacters []*Character
	World            *World
}
