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

// https://web.cs.dal.ca/~safatli/blog/?p=239
func NewSeed() uint {
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
