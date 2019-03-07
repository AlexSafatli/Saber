package main

import (
	"fmt"

	"./cmd"
	"./gen"
)

func main() {
	gen.Reseed()
	gen.InitRandomTables()
	l := gen.GenerateLanguage()
	fmt.Println(l)
	fmt.Println(l.Phonemes)
	for i := 0; i < 12; i++ {
		fmt.Println(l.Name())
	}
	cmd.Execute()
}
