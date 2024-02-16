package rng

import (
	wr "github.com/mroth/weightedrand"
)
import "golang.org/x/exp/constraints"

func ShuffleSlice(a *[]interface{}) {
	for i := range *a {
		j := rng.Intn(i + 1)
		(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
	}
}

func ShuffledSlice(a []interface{}) []interface{} {
	b := make([]interface{}, len(a))
	p := rng.Perm(len(a))
	for i, v := range p {
		b[v] = a[i]
	}
	return b
}

func Choose[S ~[]E, E constraints.Ordered](s S) E {
	i := rng.Intn(len(s))
	return s[i]
}

func ChooseWithWeights[S ~[]E, E constraints.Ordered](s S, weights []uint) E {
	var choices []wr.Choice
	for i := range s {
		choices = append(choices, wr.Choice{Item: s[i], Weight: weights[i]})
	}
	chooser, _ := wr.NewChooser(choices...)
	return chooser.Pick()
}
