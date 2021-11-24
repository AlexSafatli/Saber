package cli

import (
	"errors"
	"fmt"
	"github.com/AlexSafatli/Saber/db"
	"github.com/spf13/cobra"
)

var createRootCmd = &cobra.Command{
	Use:   "create",
	Short: "Create story elements",
	Args:  needCommandArg,
	Run:   noOpCmd,
}

var createCampaignCmd = &cobra.Command{
	Use:   "campaign <name>",
	Short: "Create a new campaign",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("need only a name for the campaign")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		client, _, err := db.Connect()
		if err != nil {
			panic(err)
		}
		cmpn, err := db.CreateCampaign(client, args[0])
		if err != nil {
			panic(err)
		}
		fmt.Printf("Generated new campaign ('%s') with ID %s.\n",
			args[0], cmpn)
	},
}
