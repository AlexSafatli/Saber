package rng

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
	GenerateConnections(w)
	return w
}

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
		w.Regions[i] = entities.Region{Name: l.Name()}
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
		r.Subregions[i] = entities.Region{Name: l.Name()}
		go PopulateRegion(&r.Subregions[i], l, complexity-1, wg)
	}
}

func GenerateConnections(w *entities.World) {
	// use seed object here
	// https://web.cs.dal.ca/~safatli/blog/?p=239
}
