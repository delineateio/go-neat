package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	c "go.delineate.io/neat/config"
	e "go.delineate.io/neat/errors"
	g "go.delineate.io/neat/git"
	u "go.delineate.io/neat/ui"
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
	Short: "Creates a new feature and performs cleanup",
	Long: `
Creates a new feature branch and performs cleanup:

* Checks out the default branch
* Pulls & fetches the default branch
* Deletes any redundnat feature branches
* Creates a new feature branch ready for development`,

	Run: func(cmd *cobra.Command, args []string) {

		newFeatureConfig(newFeature.path)

		repo := g.NewGitRepository(newFeature.path)
		branches := repo.Branches()

		branches.DefaultBranch().Checkout().Pull()

		deleteBranches(branches)
		repo.CreateBranch(newFeature.branchName).Checkout()
	},
}

func init() {
	newCmd.AddCommand(newFeatureCmd)
	newFeatureCmd.Flags().StringVarP(&newFeature.branchName, "name", "n", "", "name of the new branch to use")
	newFeatureCmd.Flags().StringVarP(&newFeature.path, "path", "p", ".", "path of the git repository")
	err := newFeatureCmd.MarkFlagRequired("name")
	e.CheckIfError(err, "failed to initalise new feature command")
}

func newFeatureConfig(path string) {
	c.MergeConfig(c.NewDefaultConfigInfo(".neat-repo", path).
		AddDefault("git.branches.default", "main").
		AddDefault("git.branches.prune", "manual"))
}

func deleteBranches(branches *g.GitBranches) {
	prune := c.GetString("git.branches.prune")
	if prune != "none" {
		names := branches.NonDefaultNames()
		if len(names) > 0 {
			if prune == "select" {
				names = u.Checklist("branches to delete", names)
			}
			branches.FilterByNames(names).Delete()
		}
		log.Info().
			Str("event", "branches_pruned").
			Int("pruned_count", len(names)).
			Send()
	}
}
