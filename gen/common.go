package gen

import (
	"math/rand"
	"time"
)

func Reseed() {
	rand.Seed(int64(time.Now().Nanosecond()))
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
