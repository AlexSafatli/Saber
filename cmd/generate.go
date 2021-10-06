package cmd

import (
	"fmt"
	"github.com/AlexSafatli/Saber/gen"
	"github.com/spf13/cobra"
)

var genRootCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates story elements",
	Run: func(cmd *cobra.Command, args []string) {
		// do stuff
	},
}

var genWorldCmd = &cobra.Command{
	Use:   "world",
	Short: "Generates a new world",
	Run: func(cmd *cobra.Command, args []string) {
		l := gen.GenerateLanguage()
		fmt.Println("Generated world: ", gen.GenerateWorld(l, 3))
	},
}

var genFamilyCmd = &cobra.Command{
	Use:   "family",
	Short: "Generates a family tree",
	Run: func(cmd *cobra.Command, args []string) {
		l := gen.GenerateLanguage()
		w := gen.GenerateWorld(l, 3)
		f := gen.GenerateFamily(l, &w.Regions[0])
		tree := gen.GenerateFamilyTree(f, w, 2)
		fmt.Print("Generated family:\n\n")
		printTreeNode(&tree.Root, 0)
	},
}

func printTreeNode(node *gen.FamilyTreeNode, level uint) {
	var tabs = ""
	var i uint = 0
	for i < level {
		tabs += "\t" // TODO string builder
		i++
	}
	fmt.Printf("%s%s (%s) with %d children\n", tabs, node.Character.Name,
		node.Character.Profession, len(node.Children))
	for _, child := range node.Children {
		printTreeNode(child, level+1)
	}
}
