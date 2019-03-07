package cmd

import (
	"../gen"
	"fmt"
	"github.com/spf13/cobra"
)

var genWorldCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a new world",
	Run: func(cmd *cobra.Command, args []string) {
		l := gen.GenerateLanguage()
		fmt.Println("Generated world: ", gen.GenerateWorld(l, 3))
	},
}
