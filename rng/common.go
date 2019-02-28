package rng

import (
	"math/rand"
	"time"
)

func Reseed() {
	rand.Seed(int64(time.Now().Nanosecond()))
}
