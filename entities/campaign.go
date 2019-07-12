package entities

type Campaign struct {
	Name       string
	DM         string
	Characters []*Character
	World      *World
}
