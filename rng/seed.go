package rng

import (
	"math/rand"
	"time"
)

var rng rand.Rand

func New() rand.Rand {
	rng.Seed(int64(time.Now().Nanosecond()))
	return rng
}
