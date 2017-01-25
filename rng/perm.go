package rng

import "math/rand"

func shuffle(a *[]interface{}) {
	for i := range *a {
		j := rand.Intn(i + 1)
		(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
	}
}

func shuffled(a []interface{}) []interface{} {
	b := make([]interface{}, len(a))
	p := rand.Perm(len(a))
	for i, v := range p {
		b[v] = a[i]
	}
	return b
}

func shuffledString(s string) string {
	a := []rune(s)
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}

func choose(a []string) string {
	i := rand.Intn(len(a))
	return a[i]
}

func chooseRune(s string) rune {
	runes := []rune(s)
	if len(runes) == 1 {
		return runes[0]
	}
	i := rand.Intn(len(runes))
	return runes[i]
}
