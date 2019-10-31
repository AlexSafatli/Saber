package main

import (
	"./cmd"
	"./gen"
)

func main() {
	gen.Reseed()
	gen.InitRandomTables()
	cmd.Execute()
}
