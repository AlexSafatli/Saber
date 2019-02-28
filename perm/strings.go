package perm

import "math/rand"

func ShuffledString(s string) string {
	a := []rune(s)
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}

func ChooseRune(s string) rune {
	runes := []rune(s)
	if len(runes) == 1 {
		return runes[0]
	}
	i := rand.Intn(len(runes))
	return runes[i]
}
