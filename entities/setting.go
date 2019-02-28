package entities

type World struct {
	Name    string
	Regions []Region
}

type Region struct {
	Name        string
	Biome       string
	Connections []*Region
	Subregions  []Region
}

type Settlement struct {
	Name      string
	Size      uint8
	Locations []Location
}

type Location struct {
	Name string
	Type uint8
}

func NewWorld(name string) *World {
	return &World{Name: name}
}
