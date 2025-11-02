package main

import (
	"os"

	cmd "github.com/jpa-rocha/adamastor/cmd"
)

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
