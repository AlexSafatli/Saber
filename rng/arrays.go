package rng

import "math/rand"
import "golang.org/x/exp/constraints"

func ShuffleSlice(a *[]interface{}) {
	for i := range *a {
		j := rand.Intn(i + 1)
		(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
	}
}

func ShuffledSlice(a []interface{}) []interface{} {
	b := make([]interface{}, len(a))
	p := rand.Perm(len(a))
	for i, v := range p {
		b[v] = a[i]
	}
	return b
}

func Choose[S ~[]E, E constraints.Ordered](s S) E {
	i := rand.Intn(len(s))
	return s[i]
}
