package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var undoCmd = &cobra.Command{
	Use:   "sync",
	Short: "A brief description of your application",
	Long:  ""}

func init() {
	neatCmd.AddCommand(undoCmd)
}
