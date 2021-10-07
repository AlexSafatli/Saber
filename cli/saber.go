package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "saber",
	Short: "Saber is a dungeon-mastering CLI for the technically competent",
	Run: func(cmd *cobra.Command, args []string) {
		// do stuff
	},
}

func Execute() {
	genRootCmd.AddCommand(genWorldCmd)
	genRootCmd.AddCommand(genFamilyCmd)
	rootCmd.AddCommand(genRootCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
