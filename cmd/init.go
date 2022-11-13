package cmd

import (
	c "github.com/delineateio/go-neat/config"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Write a config file to the home directory of current user",
	Run: func(cmd *cobra.Command, args []string) {
		c.WriteGlobalInitConfig()
	},
}

func init() {
	neatCmd.AddCommand(initCmd)
}
