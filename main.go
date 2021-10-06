package main

import (
	"github.com/AlexSafatli/Saber/cmd"
	"github.com/AlexSafatli/Saber/gen"
)

func main() {
	gen.Reseed()
	gen.InitRandomTables()
	cmd.Execute()
}
