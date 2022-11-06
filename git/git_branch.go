package git

import (
	e "github.com/delineateio/go-neat/errors"
	u "github.com/delineateio/go-neat/ui"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/rs/zerolog/log"
)

type GitBranch struct {
	RepoInfo        *GitRepository
	Name            string
	RefName         string
	Ref             *plumbing.Reference
	IsDefaultBranch bool
}

func NewBranch(repo *GitRepository, ref *plumbing.Reference) *GitBranch {
	return &GitBranch{
		RepoInfo:        repo,
		Name:            ref.Name().Short(),
		RefName:         ref.Name().String(),
		Ref:             ref,
		IsDefaultBranch: ref.Name().Short() == repo.DefaultBranch,
	}
}

func (b *GitBranch) logBranchEvent(event string) {
	log.Info().
		Str("path", b.RepoInfo.Path).
		Str("event", event).
		Str("name", b.Name).Send()
}

func (b *GitBranch) getWorktree() (*git.Worktree, git.Status) {

	worktree, err := b.RepoInfo.Repo.Worktree()
	e.CheckIfError(err, "failed to get the working tree for '%s'", b.Name)

	status, err := worktree.Status()
	e.CheckIfError(err, "failed to get the working tree status for '%s'", b.Name)

	return worktree, status
}

func (b *GitBranch) Checkout() *GitBranch {

	worktree, status := b.getWorktree()
	if !status.IsClean() {
		e.NewErr("failed to checkout '%s' because the working tree is dirty", b.Name)
	}

	err := worktree.Checkout(&git.CheckoutOptions{
		Branch: b.Ref.Name(),
	})
	e.CheckIfError(err, "failed to checkout '%s'", b.Name)

	b.logBranchEvent("branch_checked_out")
	u.Successful("succesfully checked out '%s'", b.Name)
	return b
}

func (b *GitBranch) Pull() *GitBranch {

	worktree, status := b.getWorktree()
	remotes, err := b.RepoInfo.Repo.Remotes()
	e.CheckIfError(err, "failed to retrieve the remotes for '%s'", b.Name)
	if len(remotes) == 0 {
		u.Skipped("skipped pull of '%s' as no remotes", b.Name)
		return b
	}

	if !status.IsClean() {
		e.NewErr("failed as the working tree for '%s' is dirty", b.Name)
	} else {
		err := worktree.Pull(&git.PullOptions{})
		e.CheckIfError(err, "failed to pull '%s'", b.Name)
	}

	b.logBranchEvent("branch_pulled")
	u.Successful("succesfully pulled '%s'", b.Name)
	return b
}

func (b *GitBranch) Delete() {
	err := b.RepoInfo.Repo.Storer.RemoveReference(b.Ref.Name())
	e.CheckIfError(err, "failed to delete '%s'", b.Name)

	b.logBranchEvent("branch_deleted")
	u.Successful("succesfully deleted '%s'", b.Name)
}
