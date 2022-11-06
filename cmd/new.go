package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Provides access to all 'new' commands",
}

func init() {
	neatCmd.AddCommand(newCmd)
}
