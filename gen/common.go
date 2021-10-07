package gen

import (
	"math/rand"
	"time"
)

const (
	seedReductionAmt = 40
	seedIncrementAmt = 2
)

type Possibility struct {
	// Encompasses what I refer to as a "seed object" in this post:
	// https://web.cs.dal.ca/~safatli/blog/?p=239
	value uint
}

func (p *Possibility) Reduce() {
	p.value -= seedReductionAmt
	if p.value <= 0 {
		p.value = 1
	}
}

func (p *Possibility) Increase() {
	p.value += seedIncrementAmt
	if p.value > 100 {
		p.value = 100
	}
}

func Reseed() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

func NewPossibility() *Possibility {
	return &Possibility{value: randPercentile()}
}

func randPercentile() uint {
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
