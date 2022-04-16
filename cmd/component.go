package cmd

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/spf13/cobra"
	"log"
	"rc/generate/component"
)

// componentCmd represents the component command
var componentCmd = &cobra.Command{
	Use:   "c",
	Short: "Creates a new component",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(color.InRed("No component named specified."))
		}
		createComponent(args)
	},
}

func createComponent(args []string) {
	for _, arg := range args {
		err := component.Create(arg)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func init() {
	generateCmd.AddCommand(componentCmd)
}
