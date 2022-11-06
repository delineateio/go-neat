package cmd

import (
	"os"

	c "github.com/delineateio/go-neat/config"
	l "github.com/delineateio/go-neat/logging"
	"github.com/spf13/cobra"
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
