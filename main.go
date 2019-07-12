package main

import (
	"./cmd"
	"./gen"
	"fmt"
	"github.com/spf13/viper"
)

func readConfig() {
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func main() {
	readConfig()
	gen.Reseed()
	gen.InitRandomTables()
	cmd.Execute()
}
