package rpg

import (
	"github.com/AlexSafatli/Saber/rng"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type World struct {
	Location
	Regions []Region
}

type Region struct {
	Location
	Connections []*Region `json:"-" bson:"-"`
	Subregions  []Region  `json:",omitempty" bson:",omitempty"`
}

type Settlement struct {
	Location
	Size      uint8
	Locations []Location
}

type Location struct {
	Name string
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Tags []EntityTag        `json:",omitempty" bson:",omitempty"`
	Type uint8
}

type Date struct {
	Delta                uint
	AddSpouseID          uint `json:",omitempty" bson:",omitempty"`
	RemoveSpouseID       uint `json:",omitempty" bson:",omitempty"`
	EmployerID           uint `json:",omitempty" bson:",omitempty"`
	AddCharacteristic    uint `json:",omitempty" bson:",omitempty"`
	RemoveCharacteristic uint `json:",omitempty" bson:",omitempty"`
}

const (
	LocationTypeWorld     = 0
	LocationTypeContinent = 1
	LocationTypeCountry   = 2
	LocationTypeProvince  = 3
	LocationTypeCity      = 4
	LocationTypeTown      = 5
	LocationTypeVillage   = 6
	LocationTypeHamlet    = 7
	LocationTypeThorp     = 8
	LocationTypeDistrict  = 9
	LocationTypeUnit      = 10
)

var (
	LargeLocationTypes         = []uint8{LocationTypeContinent, LocationTypeCountry, LocationTypeProvince}
	CityLocationTypes          = []uint8{LocationTypeCity, LocationTypeTown, LocationTypeVillage, LocationTypeHamlet, LocationTypeThorp}
	LocalLocationTypes         = []uint8{LocationTypeDistrict, LocationTypeUnit}
	SmallestLocationType uint8 = LocationTypeUnit
)

func (r *Region) IsLargeLocation() bool {
	return contains(r.Type, LargeLocationTypes)
}

func NewWorld(name string) *World {
	return &World{Location: Location{
		Name: name,
		Type: LocationTypeWorld,
	}}
}

func MakeRegion(name string, locType uint8) Region {
	return Region{Location: Location{
		Name: name,
		Type: locType,
	}}
}

func GenerateLargeLocationType() uint8 {
	return rng.ChooseUint8(LargeLocationTypes)
}

func GenerateCityLocationType() uint8 {
	return rng.ChooseUint8(CityLocationTypes)
}

func GenerateLocalLocationType() uint8 {
	return rng.ChooseUint8(LocalLocationTypes)
}

func NextSmallestLocationType(locType uint8) uint8 {
	if locType >= SmallestLocationType {
		return SmallestLocationType
	} else {
		return locType + 1
	}
}

func contains(v uint8, a []uint8) bool {
	for _, u := range a {
		if u == v {
			return true
		}
	}
	return false
}
