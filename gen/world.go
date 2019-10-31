package gen

import (
	"../entities"
	"sync"
)

const (
	WorldComplexitySimple     = 1
	WorldComplexityNormal     = 2
	WorldComplexityLarge      = 3
	WorldComplexityExtraLarge = 4
)

func GenerateWorld(l *Language, complexity uint8) *entities.World {
	w := entities.NewWorld(l.Name())
	PopulateWorld(w, l, complexity)
	GenerateRegionConnections(w)
	return w
}

// TODO refactor this to use different languages for different regions OR genetically change
func PopulateWorld(w *entities.World, l *Language, complexity uint8) {
	var numRegions = 0

	switch complexity {
	case WorldComplexitySimple:
		numRegions = 5
	case WorldComplexityNormal:
		numRegions = 20
	case WorldComplexityLarge:
		numRegions = 75
	case WorldComplexityExtraLarge:
		numRegions = 150
	}

	var wg sync.WaitGroup
	wg.Add(numRegions)
	w.Regions = make([]entities.Region, numRegions)
	for i := 0; i < numRegions; i++ {
		w.Regions[i] = entities.MakeRegion(l.Name(), entities.GenerateLargeLocationType())
		go PopulateRegion(&w.Regions[i], l, complexity, &wg)
	}
	wg.Wait()
}

func PopulateRegion(r *entities.Region, l *Language, complexity uint8, wg *sync.WaitGroup) {
	defer wg.Done()
	if complexity == 0 {
		return
	}
	var numRegions = 0

	switch complexity {
	case WorldComplexitySimple:
		numRegions = 3
	case WorldComplexityNormal:
		numRegions = 7
	case WorldComplexityLarge:
		numRegions = 16
	case WorldComplexityExtraLarge:
		numRegions = 64
	}

	wg.Add(numRegions)
	r.Subregions = make([]entities.Region, numRegions)
	for i := 0; i < numRegions; i++ {
		var locationType uint8
		if r.IsLargeLocation() {
			locationType = entities.GenerateCityLocationType()
		} else {
			locationType = entities.NextSmallestLocationType(r.Type)
		}
		r.Subregions[i] = entities.MakeRegion(l.Name(), locationType)
		go PopulateRegion(&r.Subregions[i], l, complexity-1, wg)
	}
}

// https://web.cs.dal.ca/~safatli/blog/?p=239
func GenerateRegionConnections(w *entities.World) {
	var wg sync.WaitGroup
	wg.Add(len(w.Regions))
	s := NewSeed()
	for i := 0; i < len(w.Regions); i++ {
		for j := i + 1; j < len(w.Regions); j++ {
			var r uint
			r = NewSeed()
			if r < s {
				s -= SeedReductionAmt
				if s <= 0 {
					s = 1
				}
				makeRegionConnection(w, i, j)
			} else {
				s += SeedIncrementAmt
				if s > 100 {
					s = 100
				}
			}
		}
		go GenerateSubregionConnections(&w.Regions[i], &wg)
	}
}

func GenerateSubregionConnections(re *entities.Region, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(len(re.Subregions))
	s := NewSeed()
	for i := 0; i < len(re.Subregions); i++ {
		for j := i + 1; j < len(re.Subregions); j++ {
			var r uint
			r = NewSeed()
			if r < s {
				s -= SeedReductionAmt
				if s <= 0 {
					s = 1
				}
				makeSubregionConnection(re, i, j)
			} else {
				s += SeedIncrementAmt
				if s > 100 {
					s = 100
				}
			}
		}
		go GenerateSubregionConnections(&re.Subregions[i], wg)
	}
}

func makeRegionConnection(w *entities.World, i, j int) {
	w.Regions[i].Connections = append(w.Regions[i].Connections, &w.Regions[j])
	w.Regions[j].Connections = append(w.Regions[j].Connections, &w.Regions[i])
}

func makeSubregionConnection(re *entities.Region, i, j int) {
	re.Subregions[i].Connections = append(re.Subregions[i].Connections, &re.Subregions[j])
	re.Subregions[j].Connections = append(re.Subregions[j].Connections, &re.Subregions[i])
}
