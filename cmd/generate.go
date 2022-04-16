package cmd

import (
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "g",
	Short: "Generates and/or modifies files based on a schema/command.",
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
