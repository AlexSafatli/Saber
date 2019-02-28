package main

import (
	"fmt"

	"./cmd"
	"./rng"
)

func main() {
	rng.Seed()
	l := rng.NewLanguage()
	fmt.Println(l)
	fmt.Println(l.Phonemes)
	for i := 0; i < 12; i++ {
		fmt.Println(l.Name())
	}
	cmd.Execute()
}