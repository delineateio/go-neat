package cmd

import (
	"os"

	"github.com/spf13/cobra"
	c "go.delineate.io/neat/config"
	l "go.delineate.io/neat/logging"
)

var neatCmd = &cobra.Command{
	Use:   "neat",
	Short: "neat is an opinionated set of development commands",
}

func Execute() {
	c.ReadGlobalConfig()
	l.InitialiseLogging()
	err := neatCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
