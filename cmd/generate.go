package cmd

import (
	"../rng"
	"fmt"
	"github.com/spf13/cobra"
)

var genWorldCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a new world",
	Run: func(cmd *cobra.Command, args []string) {
		l := rng.GenerateLanguage()
		fmt.Println("Generated world: ", rng.GenerateWorld(l, 3))
	},
}
