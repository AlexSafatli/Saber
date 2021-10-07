package main

import (
	"github.com/AlexSafatli/Saber/cli"
	"github.com/AlexSafatli/Saber/gen"
)

func main() {
	gen.Reseed()
	gen.InitRandomTables()
	cli.Execute()
}
