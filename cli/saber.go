package cli

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "saber",
	Short: "Saber is a dungeon-mastering CLI for the technically competent",
	Args:  needCommandArg,
	Run:   noOpCmd,
}

func needCommandArg(_ *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("need a command to run")
	}
	return nil
}

func noOpCmd(_ *cobra.Command, _ []string) {

}

func Execute() {
	// gen
	genRootCmd.AddCommand(genWorldCmd)
	genRootCmd.AddCommand(genFamilyCmd)
	genRootCmd.Flags().Uint8VarP(&complexityFlag, "complexity", "c", 1, "a complexity for the story element being generated")

	// root
	rootCmd.AddCommand(genRootCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
