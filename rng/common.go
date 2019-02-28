package rng

import (
	"math/rand"
	"time"
)

func Seed() {
	rand.Seed(int64(time.Now().Nanosecond()))
}
