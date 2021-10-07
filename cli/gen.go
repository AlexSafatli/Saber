package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/AlexSafatli/Saber/gen"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var complexityFlag uint8

var genRootCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate story elements",
	Args:  needCommandArg,
	Run:   noOpCmd,
}

var genWorldCmd = &cobra.Command{
	Use:   "world <output_json>",
	Short: "Generate a new world",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("need only a path for the output JSON")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		l := gen.GenerateLanguage()
		w := gen.GenerateWorld(l, complexityFlag)
		file, _ := json.MarshalIndent(w, "", " ")
		if err := ioutil.WriteFile(args[0], file, 0644); err != nil {
			panic(err)
		}
		fmt.Printf("Generated world (complexity %d) and saved to %s.\n",
			complexityFlag, args[0])
	},
}

var genFamilyCmd = &cobra.Command{
	Use:   "family <output_json>",
	Short: "Generate a family tree",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("need only a path for the output JSON")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		l := gen.GenerateLanguage()
		w := gen.GenerateWorld(l, complexityFlag)
		f := gen.GenerateFamily(l, &w.Regions[0])
		tree := gen.GenerateFamilyTree(f, w, 3)
		file, _ := json.MarshalIndent(tree, "", " ")
		if err := ioutil.WriteFile(args[0], file, 0644); err != nil {
			panic(err)
		}
		fmt.Printf("Generated family and saved to %s.\n", args[0])
	},
}
