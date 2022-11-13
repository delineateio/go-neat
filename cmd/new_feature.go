package cmd

import (
	c "github.com/delineateio/go-neat/config"
	e "github.com/delineateio/go-neat/errors"
	g "github.com/delineateio/go-neat/git"
	u "github.com/delineateio/go-neat/ui"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var new_feature_name = ""
var new_feature_path = "."

var newFeatureCmd = &cobra.Command{
	Use:   "feature",
	Short: "Creates a new feature and performs clean up",
	Run: func(cmd *cobra.Command, args []string) {

		config := getConfig(new_feature_path)

		repo := g.NewGitRepository(new_feature_path)
		if !repo.Exists {
			e.NewErr("'%s' is not a repository", new_feature_path)
		}

		branches := repo.Branches()
		branches.DefaultBranch().Checkout().Pull()

		deleteBranches(branches, config)
		repo.CreateBranch(new_feature_name).Checkout()
	},
}

func init() {
	newCmd.AddCommand(newFeatureCmd)
	addStrFlag(newFeatureCmd, "name", "name of the new feature", &new_feature_name)
	addStrFlag(newFeatureCmd, "path", "path of the git repository", &new_feature_path)
	addRequired("name")
}

func getConfig(path string) c.NeatConfig {

	i := c.NewDefaultConfig(".neat", path).
		AddDefault("git.branches.prune", "auto")

	return c.GetConfig(i)
}

func deleteBranches(branches *g.GitBranches, config c.NeatConfig) {

	prune := config.GetString("branches.prune")
	names := branches.NonDefaultNames()

	switch prune {
	case "none":
		names = []string{}
	case "select":
		if len(names) > 0 {
			names = u.Checklist("branches to delete", names)
		} else {
			u.Skipped("skipped selecting branches non-main branches")
		}
	case "auto":
	default:
		e.NewErr("unexpected value for pruning option '%s'", prune)
	}

	branches.FilterByNames(names).Delete()
	log.Info().
		Str("event", "branches_pruned").
		Int("pruned_count", len(names)).
		Send()
}
