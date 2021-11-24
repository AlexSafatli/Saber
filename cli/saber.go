package cli

import (
	"context"
	"errors"
	"fmt"
	"github.com/AlexSafatli/Saber/db"
	"github.com/AlexSafatli/Saber/rpg"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/mongo"
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

func getCampaign(query string) (client *mongo.Client, ctx context.Context, c *rpg.Campaign) {
	var err error
	client, ctx, err = db.Connect()
	if err != nil {
		panic(err)
	}
	campaigns := db.Campaigns(client)
	err = campaigns.FindOne(ctx, rpg.Campaign{Name: query}).Decode(c)
	if err != nil {
		panic(err)
	}
	return
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
