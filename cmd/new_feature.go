package cmd

import (
	c "github.com/delineateio/go-neat/config"
	e "github.com/delineateio/go-neat/errors"
	g "github.com/delineateio/go-neat/git"
	u "github.com/delineateio/go-neat/ui"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type newFeatureContext struct {
	branchName string
	path       string
}

var newFeature = newFeatureNewContext()

func newFeatureNewContext() newFeatureContext {
	return newFeatureContext{}
}

var newFeatureCmd = &cobra.Command{
	Use:   "feature",
	Short: "Creates a new feature and performs clean up",
	Long: `
Creates a new feature branch and performs clean up:

* Checks out the default branch
* Pulls & fetches the default branch
* Deletes any redundant feature branches
* Creates a new feature branch ready for development`,

	Run: func(cmd *cobra.Command, args []string) {

		newFeatureConfig(newFeature.path)

		repo := g.NewGitRepository(newFeature.path)
		if !repo.Exists {
			e.NewErr("'%s' is not a repository", newFeature.path)
		}

		branches := repo.Branches()
		branches.DefaultBranch().Checkout().Pull()

		deleteBranches(branches)
		repo.CreateBranch(newFeature.branchName).Checkout()
	},
}

func init() {
	newCmd.AddCommand(newFeatureCmd)
	newFeatureCmd.Flags().StringVar(&newFeature.branchName, "name", "", "name of the new branch to use")
	newFeatureCmd.Flags().StringVar(&newFeature.path, "path", ".", "path of the git repository")
	err := newFeatureCmd.MarkFlagRequired("name")
	e.CheckIfError(err, "failed to initialise new feature command")
}

func newFeatureConfig(path string) {

	c.NewDefaultConfig(".neat", path).
		AddDefault("git.branches.prune", "auto").
		Merge()
}

func deleteBranches(branches *g.GitBranches) {

	prune := c.GetString("git.branches.prune")
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
