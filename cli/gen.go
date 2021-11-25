package cli

import (
	"context"
	"errors"
	"fmt"
	"github.com/AlexSafatli/Saber/db"
	"github.com/AlexSafatli/Saber/gen"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var complexityFlag uint8

var genRootCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate story elements",
	Args:  needCommandArg,
	Run:   noOpCmd,
}

var genWorldCmd = &cobra.Command{
	Use:   "world <campaign>",
	Short: "Generate a new world",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("need only a name or ID for the campaign")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		client, ctx, c := getCampaign(args[0])
		cmpn := db.Campaign(client, c.ID.Hex())
		l := gen.GenerateLanguage()
		w := gen.GenerateWorld(l, complexityFlag)
		_, err := cmpn.Collection("languages").InsertOne(ctx, l)
		if err != nil {
			panic(err)
		}
		_, err = cmpn.Collection("worlds").InsertOne(ctx, w)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Generated world (complexity %d, name '%s') for %s.\n",
			complexityFlag, args[0], w.Name)
	},
}

var genFamilyCmd = &cobra.Command{
	Use:   "family <campaign>",
	Short: "Generate a family tree",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("need a name or ID for the campaign")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		client, ctx, c := getCampaign(args[0])
		cmpn := db.Campaign(client, c.ID.Hex())
		l := gen.GenerateLanguage()
		w := gen.GenerateWorld(l, complexityFlag)
		f := gen.GenerateFamily(l, &w.Regions[0])
		tree := gen.GenerateFamilyTree(f, w, 3)
		_, err := cmpn.Collection("languages").InsertOne(ctx, l)
		if err != nil {
			panic(err)
		}
		_, err = cmpn.Collection("worlds").InsertOne(ctx, w)
		if err != nil {
			panic(err)
		}
		insertFamilyTreeNode(cmpn, ctx, &tree.Root)
		fmt.Printf("Generated family '%s' and world '%s' for %s.\n", f.Surname, w.Name, args[0])
	},
}

func insertFamilyTreeNode(cmpn *mongo.Database, ctx context.Context, node *gen.FamilyTreeNode) {
	if node.Character.ID != primitive.NilObjectID {
		return
	}
	node.Character.MotherID = node.Mother.Character.ID
	node.Character.FatherID = node.Father.Character.ID
	res, err := cmpn.Collection("characters").InsertOne(ctx, node.Character)
	if err != nil {
		panic(err)
	}
	node.Character.ID = res.InsertedID.(primitive.ObjectID)
	insertFamilyTreeNode(cmpn, ctx, node.Spouse)
	for i := range node.Children {
		insertFamilyTreeNode(cmpn, ctx, node.Children[i])
	}
}
