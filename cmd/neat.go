package cmd

import (
	"os"

	l "github.com/delineateio/go-neat/logging"
	"github.com/spf13/cobra"
)

const DOT_DIR = "."

var neatCmd = &cobra.Command{
	Use:   "neat",
	Short: "neat is an opinionated set of development commands",
}

func Execute() {
	l.InitialiseLogging()
	err := neatCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
