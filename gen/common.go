package gen

import (
	"math/rand"
	"time"
)

const (
	SeedReductionAmt = 40
	SeedIncrementAmt = 2
)

func Reseed() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

func NewSeed() uint { // Seed object; see docs
	return uint(rand.Int63n(100) << 1 >> 1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
