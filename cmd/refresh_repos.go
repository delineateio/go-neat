package cmd

import (
	"os"
	"path/filepath"

	e "github.com/delineateio/go-neat/errors"
	g "github.com/delineateio/go-neat/git"
	u "github.com/delineateio/go-neat/ui"
	"github.com/spf13/cobra"
)

type newRefreshRepos struct {
	path string
}

var refreshRepos = refreshReposContext()

func refreshReposContext() newRefreshRepos {
	return newRefreshRepos{}
}

type org struct {
	dir  *os.DirEntry
	path string
}

func (o *org) repos() *[]g.GitRepository {

	repos := make([]g.GitRepository, 0)
	items, err := os.ReadDir(o.path)
	e.CheckIfError(err, "path not found '%s'", refreshRepos.path)
	for _, item := range items {
		if item.IsDir() {
			path := filepath.Join(o.path, item.Name())
			repos = append(repos, *g.NewGitRepository(path))
		}
	}
	return &repos
}

func newOrg(orgDir os.DirEntry) *org {

	relPath := filepath.Join(refreshRepos.path, orgDir.Name())
	absPath, err := filepath.Abs(relPath)
	e.CheckIfError(err, "can't access path '%s'", relPath)

	return &org{
		dir:  &orgDir,
		path: absPath,
	}
}

var refreshReposCmd = &cobra.Command{
	Use:   "repos",
	Short: "Creates a new feature and performs clean up",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		for _, org := range getOrgs() {
			for _, repo := range *org.repos() {
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

func getOrgs() []org {

	orgs := make([]org, 0)

	root, _ := filepath.Abs(refreshRepos.path)
	items, err := os.ReadDir(root)
	e.CheckIfError(err, "path not found '%s'", refreshRepos.path)
	for _, item := range items {
		if item.IsDir() {
			orgs = append(orgs, *newOrg(item))
		}
	}
	return orgs
}

func init() {
	refreshCmd.AddCommand(refreshReposCmd)
	newCmd.AddCommand(newFeatureCmd)
	refreshReposCmd.Flags().StringVar(&refreshRepos.path, "path", ".", "path of repositories")
}
