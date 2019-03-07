package rng

import "math/rand"

func Shuffle(a *[]interface{}) {
	for i := range *a {
		j := rand.Intn(i + 1)
		(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
	}
}

func Shuffled(a []interface{}) []interface{} {
	b := make([]interface{}, len(a))
	p := rand.Perm(len(a))
	for i, v := range p {
		b[v] = a[i]
	}
	return b
}

func Choose(a []string) string {
	i := rand.Intn(len(a))
	return a[i]
}

func ChooseUint8(a []uint8) uint8 {
	i := rand.Intn(len(a))
	return a[i]
}
