package cmd

import (
	c "github.com/delineateio/go-neat/config"
	"github.com/spf13/cobra"
)

var initRepoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Provides access to all 'new' commands",

	Run: func(cmd *cobra.Command, args []string) {
		c.WriteRepoInitConfig(".")
	},
}

func init() {
	initCmd.AddCommand(initRepoCmd)
	initCmd.Flags().StringVar(&newFeature.path, "path", ".", "path of the git repository")
}
