//nolint:gochecknoinits,gochecknoglobals // Cobra sub commands need to be initiated and are called by main command.
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jpa-rocha/adamastor/internal/server"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve your files",
	RunE:  server.RunServe,
}

func init() {
	RootCmd.AddCommand(ServeCmd)
}
