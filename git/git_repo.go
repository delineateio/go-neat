package git

import (
	"path/filepath"
	"strings"

	c "github.com/delineateio/go-neat/config"
	e "github.com/delineateio/go-neat/errors"
	u "github.com/delineateio/go-neat/ui"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/rs/zerolog/log"
)

type GitRepository struct {
	Name          string
	Path          string
	AbsPath       string
	Exists        bool
	Repo          *git.Repository
	DefaultBranch string
}

func NewGitRepository(path string) *GitRepository {

	absPath, err := filepath.Abs(path)
	e.CheckIfError(err, "failed to get the absolute path for '%s'", path)

	repo, err := git.PlainOpen(absPath)
	if err != nil {
		log.Warn().Msg(err.Error())
	}

	info := &GitRepository{
		Name:          strings.ToLower(filepath.Base(path)),
		Path:          path,
		AbsPath:       absPath,
		Exists:        err == nil,
		Repo:          repo,
		DefaultBranch: c.GetString("git.branches.default"),
	}
	info.log()
	return info
}

func (r *GitRepository) Branches() *GitBranches {
	branches := NewBranches()
	iter, err := r.Repo.Branches()
	e.CheckIfError(err, "failed to get the branches the repo")

	err = iter.ForEach(func(ref *plumbing.Reference) error {
		if ref.Name().IsBranch() {
			branch := NewBranch(r, ref)
			branches[branch.Name] = *branch
		}
		return nil
	})
	e.CheckIfError(err, "failed to load the branches the repo")
	return &branches
}

func (r *GitRepository) HasBranch(branchName string) (*GitBranch, bool) {

	branchName = normalizeBranchName(branchName)
	for _, v := range *r.Branches() {
		if v.Name == branchName {
			return &v, true
		}
	}
	return nil, false
}

func normalizeBranchName(branchName string) string {
	branchName = strings.ToLower(branchName)
	return strings.ReplaceAll(branchName, " ", "")
}

func (r *GitRepository) CreateBranch(branchName string) *GitBranch {

	branchName = normalizeBranchName(branchName)

	b, ok := r.HasBranch(branchName)
	if ok {
		u.Skipped("skipped creation of '%s' as it exists", b.Name)
		return b
	}

	headRef, err := r.Repo.Head()
	e.CheckIfError(err, "failed to retrieve the HEAD")

	branchRef := plumbing.NewBranchReferenceName(branchName)
	ref := plumbing.NewHashReference(branchRef, headRef.Hash())

	err = r.Repo.Storer.SetReference(ref)
	e.CheckIfError(err, "failed to create new branch '%s'", branchName)

	branch := r.Branches().GetByName(branchName)
	u.Successful("successfully created '%s' on '%s'",
		branch.Name,
		branch.Info.Name)
	return branch
}

func (r *GitRepository) log() {
	log.Debug().
		Str("event", "repo_located").
		Str("path", r.AbsPath).
		Send()
}
