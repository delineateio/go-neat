package cmd

import (
	g "github.com/delineateio/go-neat/git"
	u "github.com/delineateio/go-neat/ui"
	"github.com/spf13/cobra"
)

var refresh_repos_path = DOT_DIR

var refreshReposCmd = &cobra.Command{
	Use:   "repos",
	Short: "Refreshes the repos in sub directories of the root path",
	Run: func(cmd *cobra.Command, args []string) {
		for _, org := range *g.GetLocalGitOrgs(refresh_repos_path) {
			for _, repo := range *org.Repos() {
				if repo.Exists {
					repo.Branches().
						DefaultBranch().
						Checkout().
						Pull()
				} else {
					u.Skipped("skipped directory '%s' as not a git repo", repo.Path)
				}
			}
		}
	},
}

func init() {
	refreshCmd.AddCommand(refreshReposCmd)
	addStrFlag(refreshReposCmd, "path", "path of repositories", &refresh_repos_path)
}
