package cmd

import (
	c "github.com/delineateio/go-neat/config"
	"github.com/spf13/cobra"
)

var init_repo_path = DOT_DIR

var initRepoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Write a config file to the root of a repo",
	Run: func(cmd *cobra.Command, args []string) {
		c.WriteRepoInitConfig(init_repo_path)
	},
}

func init() {
	initCmd.AddCommand(initRepoCmd)
	addStrFlag(initRepoCmd, "path", "path of the git repository", &init_repo_path)
}
