package cmd

import (
	"github.com/spf13/cobra"
	"rc/generate/component"
	"rc/pkg/fmtutils"
)

var isReturnComponentSet bool

// componentCmd represents the component command
var componentCmd = &cobra.Command{
	Use:   "c",
	Short: "Creates a new component",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmtutils.FatallnMsg("No component name specified.")
		}
		createComponent(args, isReturnComponentSet)
	},
}

func createComponent(args []string, isReturnComponentSet bool) {
	for _, arg := range args {
		err := component.Create(arg, isReturnComponentSet)
		if err != nil {
			fmtutils.Fatalln(err)
		}
	}
}

func init() {
	generateCmd.AddCommand(componentCmd)
	componentCmd.Flags().BoolVarP(&isReturnComponentSet, "return", "r", false,
		"Creates the component with a return structure.")
}
