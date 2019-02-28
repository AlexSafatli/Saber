package rng

import "math/rand"

func NewSeed() uint {
	return uint(rand.Int63n(100) << 1 >> 1)
}
