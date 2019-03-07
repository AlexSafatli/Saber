package gen

import "math/rand"

const (
	SeedReductionAmt = 40
	SeedIncrementAmt = 2
)

func NewSeed() uint {
	return uint(rand.Int63n(100) << 1 >> 1)
}
