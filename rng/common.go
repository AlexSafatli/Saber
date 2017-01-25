package rng

import (
  "time"
  "math/rand"
)

func Seed() {
  rand.Seed(int64(time.Now().Nanosecond()))
}