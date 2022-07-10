package cmd

import (
	"github.com/spf13/cobra"
	"rc/generate/component"
	"rc/pkg/fmtutils"
)

// componentCmd represents the component command
var componentCmd = &cobra.Command{
	Use:   "c",
	Short: "Creates a new component",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmtutils.FatallnMsg("No component name specified.")
		}
		createComponent(args)
	},
}

func createComponent(args []string) {
	for _, arg := range args {
		err := component.Create(arg)
		if err != nil {
			fmtutils.Fatalln(err)
		}
	}
}

func init() {
	generateCmd.AddCommand(componentCmd)
}
