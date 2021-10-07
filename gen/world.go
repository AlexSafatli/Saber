package gen

import (
	"github.com/AlexSafatli/Saber/rpg"
	"sync"
)

const (
	WorldComplexitySimple     = 1
	WorldComplexityNormal     = 2
	WorldComplexityLarge      = 3
	WorldComplexityExtraLarge = 4
)

func GenerateWorld(l *Language, complexity uint8) *rpg.World {
	w := rpg.NewWorld(l.Name())
	PopulateWorld(w, l, complexity)
	GenerateRegionConnections(w)
	return w
}

// TODO refactor this to use different languages for different regions OR genetically change
func PopulateWorld(w *rpg.World, l *Language, complexity uint8) {
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
	w.Regions = make([]rpg.Region, numRegions)
	for i := 0; i < numRegions; i++ {
		w.Regions[i] = rpg.MakeRegion(l.Name(), rpg.GenerateLargeLocationType())
		go PopulateRegion(&w.Regions[i], l, complexity, &wg)
	}
	wg.Wait()
}

func PopulateRegion(r *rpg.Region, l *Language, complexity uint8, wg *sync.WaitGroup) {
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
	r.Subregions = make([]rpg.Region, numRegions)
	for i := 0; i < numRegions; i++ {
		var locationType uint8
		if r.IsLargeLocation() {
			locationType = rpg.GenerateCityLocationType()
		} else {
			locationType = rpg.NextSmallestLocationType(r.Type)
		}
		r.Subregions[i] = rpg.MakeRegion(l.Name(), locationType)
		go PopulateRegion(&r.Subregions[i], l, complexity-1, wg)
	}
}

func GenerateRegionConnections(w *rpg.World) {
	var wg sync.WaitGroup
	wg.Add(len(w.Regions))
	s := NewPossibility()
	for i := 0; i < len(w.Regions); i++ {
		for j := i + 1; j < len(w.Regions); j++ {
			var r uint
			r = randPercentile()
			if r < s.value {
				s.Reduce()
				makeRegionConnection(w, i, j)
			} else {
				s.Increase()
			}
		}
		go GenerateSubregionConnections(&w.Regions[i], &wg)
	}
}

func GenerateSubregionConnections(re *rpg.Region, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(len(re.Subregions))
	s := NewPossibility()
	for i := 0; i < len(re.Subregions); i++ {
		for j := i + 1; j < len(re.Subregions); j++ {
			var r uint
			r = randPercentile()
			if r < s.value {
				s.Reduce()
				makeSubregionConnection(re, i, j)
			} else {
				s.Increase()
			}
		}
		go GenerateSubregionConnections(&re.Subregions[i], wg)
	}
}

func makeRegionConnection(w *rpg.World, i, j int) {
	w.Regions[i].Connections = append(w.Regions[i].Connections, &w.Regions[j])
	w.Regions[j].Connections = append(w.Regions[j].Connections, &w.Regions[i])
}

func makeSubregionConnection(re *rpg.Region, i, j int) {
	re.Subregions[i].Connections = append(re.Subregions[i].Connections, &re.Subregions[j])
	re.Subregions[j].Connections = append(re.Subregions[j].Connections, &re.Subregions[i])
}
