package cmd

import (
	"fmt"
	"github.com/mouuff/go-rocket-update/pkg/provider"
	"github.com/mouuff/go-rocket-update/pkg/updater"
	"github.com/spf13/cobra"
	"os"
	"rc/pkg/fmtutils"
	"runtime"
)

var AppVersion = "0.0.2"
var showVersion bool
var tryUpdate bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "rc",
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
			handleShowVersion()
		}

		if tryUpdate {
			handleUpdate()
		}

		_ = cmd.Help()
	},
}

func handleShowVersion() {
	defer os.Exit(0)
	fmt.Println("Current version: " + AppVersion)
}

func handleUpdate() {
	defer os.Exit(0)
	u := &updater.Updater{
		Provider: &provider.Github{
			RepositoryURL: "https://github.com/narvikd/rc",
			ArchiveName:   getArchiveName(),
		},
		ExecutableName: fmt.Sprintf("rc"),
		Version:        AppVersion,
	}

	fmt.Println("Current version: " + u.Version)
	fmt.Println("Looking for updates...")

	status, err := u.Update()
	if err != nil {
		fmtutils.FatallnMsg("Couldn't update. Error: " + err.Error())
	}

	if status == updater.Updated {
		fmt.Println("RC found a new version and updated itself, please restart.")
	}
}

func getArchiveName() string {
	if runtime.GOOS == "windows" {
		return fmt.Sprintf("rc_%s-%s.zip", runtime.GOOS, runtime.GOARCH)
	}
	return fmt.Sprintf("rc_%s-%s.tar.gz", runtime.GOOS, runtime.GOARCH)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.Flags().BoolVarP(&showVersion, "version", "v", false,
		"Shows current rc version.")
	rootCmd.Flags().BoolVarP(&tryUpdate, "update", "u", false,
		"Updates rc version to the latest version.")
}
