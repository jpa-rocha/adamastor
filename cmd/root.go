//nolint:gochecknoglobals // Cobra commands need to be accessible to main in order to be executed.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "github.com/jpa-rocha/adamastor",
	Short: "Go Web Server",
	Long:  "Very Large and Scary Web Server",

	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
