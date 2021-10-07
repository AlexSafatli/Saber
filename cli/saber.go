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
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("need a command to run")
		}
		return nil
	},
	Run: noOpCmd,
}

func noOpCmd(_ *cobra.Command, _ []string) {

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
