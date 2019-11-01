package rng

import "math/rand"

const MaxSmallNumber = 5

func RandomIndex(end int) int {
	return rand.Intn(end)
}

func RandomSmallNumber() int {
	return rand.Intn(MaxSmallNumber)
}

func RandomBoolean() bool {
	return rand.Intn(2) == 1
}
