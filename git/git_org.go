package git

import (
	"os"
	"path/filepath"

	e "github.com/delineateio/go-neat/errors"
)

type GitOrg struct {
	DirEntry *os.DirEntry
	Path     string
}

func NewGitOrg(root string, orgDir os.DirEntry) *GitOrg {
	path := filepath.Join(root, orgDir.Name())
	absPath, err := filepath.Abs(path)
	e.CheckIfError(err, "failed to get the absolute path for '%s'", path)
	return &GitOrg{
		DirEntry: &orgDir,
		Path:     absPath,
	}
}

func (o *GitOrg) Repos() *[]GitRepository {
	repos := make([]GitRepository, 0)
	items, err := os.ReadDir(o.Path)
	e.CheckIfError(err, "path not found '%s'", o.Path)
	for _, item := range items {
		if item.IsDir() {
			path := filepath.Join(o.Path, item.Name())
			repos = append(repos, *NewGitRepository(path))
		}
	}
	return &repos
}

func GetLocalGitOrgs(path string) *[]GitOrg {

	orgs := make([]GitOrg, 0)
	root, _ := filepath.Abs(path)
	items, err := os.ReadDir(root)
	e.CheckIfError(err, "path not found '%s'", path)
	for _, item := range items {
		if item.IsDir() {
			orgs = append(orgs, *NewGitOrg(root, item))
		}
	}
	return &orgs
}
