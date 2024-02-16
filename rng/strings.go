package rng

func ShuffledString(s string) string {
	a := []rune(s)
	for i := range a {
		j := rng.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}

func ChooseRune(s string) rune {
	runes := []rune(s)
	if len(runes) == 1 {
		return runes[0]
	}
	i := rng.Intn(len(runes))
	return runes[i]
}
