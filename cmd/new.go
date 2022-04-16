package cmd

import (
	"rc/pkg/fmtutils"
	"rc/pkg/osutils"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new React project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmtutils.FatallnMsg("No project name specified.")
		}
		createNewProject(args)
	},
}

func createNewProject(args []string) {
	_, err := osutils.NewProcess(false,
		"npx", "create-react-app", args[0], "--template", "typescript")
	if err != nil {
		fmtutils.Fatalln(err)
	}
}

func init() {
	rootCmd.AddCommand(newCmd)
}
