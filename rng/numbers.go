package rng

const MaxSmallNumber = 5

func RandomIndex(end int) int {
	return rng.Intn(end)
}

func RandomSmallNumber() int {
	return rng.Intn(MaxSmallNumber)
}

func RandomBoolean() bool {
	return rng.Intn(2) == 1
}
