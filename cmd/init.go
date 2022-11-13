package cmd

import (
	c "github.com/delineateio/go-neat/config"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Provides access to all 'new' commands",

	Run: func(cmd *cobra.Command, args []string) {

		c.WriteGlobalInitConfig()
	},
}

func init() {
	neatCmd.AddCommand(initCmd)
}
