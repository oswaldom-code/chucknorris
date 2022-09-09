package api

import (
	"os"

	"github.com/oswaldom-code/chucknorris/pkg/config"
	"github.com/spf13/cobra"
)

func init() {
	config.Load("api")
}

var RootCmd = &cobra.Command{
	Use:   "presidium",
	Short: "Presidium Enterprise Backend Services",
	Long: `
		This application is the main entry point to the backend services that are needed to run Presidium
		Enterprise.`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
