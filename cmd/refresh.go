package cmd

import (
	"github.com/spf13/cobra"
)

var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Provides access to all 'new' commands",
}

func init() {
	neatCmd.AddCommand(refreshCmd)
}
