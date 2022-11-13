package git

import (
	e "github.com/delineateio/go-neat/errors"
	u "github.com/delineateio/go-neat/ui"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/rs/zerolog/log"
)

var default_branches = []string{"master", "main", "trunk"}

type GitBranch struct {
	Info    *GitRepository
	Name    string
	RefName string
	Ref     *plumbing.Reference
}

func NewBranch(repo *GitRepository, ref *plumbing.Reference) *GitBranch {
	return &GitBranch{
		Info:    repo,
		Name:    ref.Name().Short(),
		RefName: ref.Name().String(),
		Ref:     ref,
	}
}

func (b *GitBranch) IsDefault() bool {

	for _, default_branch := range default_branches {
		if b.Name == default_branch {
			return true
		}
	}
	return false
}

func (b *GitBranch) logBranchEvent(event string) {
	log.Info().
		Str("path", b.Info.Path).
		Str("event", event).
		Str("name", b.Name).Send()
}

func (b *GitBranch) getWorktree() (*git.Worktree, git.Status) {

	worktree, err := b.Info.Repo.Worktree()
	e.CheckIfError(err, "failed to get the working tree for '%s'", b.Name)

	status, err := worktree.Status()
	e.CheckIfError(err, "failed to get working tree status for '%s' of '%s'",
		b.Name,
		b.Info.Name)

	return worktree, status
}

func (b *GitBranch) Checkout() *GitBranch {

	worktree, status := b.getWorktree()
	if !status.IsClean() {
		e.NewErr("failed to checkout '%s' of '%s' because working tree is dirty",
			b.Name,
			b.Info.Name)
	}

	err := worktree.Checkout(&git.CheckoutOptions{
		Branch: b.Ref.Name(),
	})
	e.CheckIfError(err, "failed to checkout '%s'", b.Name)

	b.logBranchEvent("branch_checked_out")
	u.Successful("successfully checked out '%s' on '%s'",
		b.Name,
		b.Info.Name)
	return b
}

func (b *GitBranch) Pull() *GitBranch {

	worktree, status := b.getWorktree()
	remotes, err := b.Info.Repo.Remotes()
	e.CheckIfError(err, "failed to retrieve the remotes for '%s' on '%s'",
		b.Name,
		b.Info.Name)
	if len(remotes) == 0 {
		u.Skipped("skipped pull of '%s' on '%s' as no remotes",
			b.Name,
			b.Info.Name)
		return b
	}

	if !status.IsClean() {
		e.NewErr("failed as working tree for '%s' on '%s' is dirty",
			b.Name,
			b.Info.Name)
	} else {
		err := worktree.Pull(&git.PullOptions{})
		if err.Error() == "already up-to-date" {
			u.Skipped("skipped pull of '%s' on '%s' as already up to date",
				b.Name,
				b.Info.Name)
			return b
		} else {
			e.CheckIfError(err, "failed to pull '%s'", b.Name)
		}
	}

	b.logBranchEvent("branch_pulled")
	u.Successful("successfully pulled '%s'", b.Name)
	return b
}

func (b *GitBranch) Delete() {
	err := b.Info.Repo.Storer.RemoveReference(b.Ref.Name())
	e.CheckIfError(err, "failed to delete '%s'", b.Name)

	b.logBranchEvent("branch_deleted")
	u.Successful("successfully deleted '%s' from '%s'",
		b.Name,
		b.Info.Name)
}
