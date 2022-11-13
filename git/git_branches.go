package git

import (
	e "github.com/delineateio/go-neat/errors"
)

type GitBranches map[string]GitBranch

func NewBranches() GitBranches {
	return make(GitBranches)
}

func (b *GitBranches) NonDefaultNames() []string {
	names := make([]string, 0, len(*b))
	for name, branch := range *b {
		if !branch.IsDefault() {
			names = append(names, name)
		}
	}
	return names
}

func (b *GitBranches) FilterByNames(names []string) *GitBranches {
	set := make(map[string]bool)
	for _, name := range names {
		set[name] = true
	}
	for name := range *b {
		if _, exists := set[name]; !exists {
			delete(*b, name)
		}
	}
	return b
}

func (b *GitBranches) Delete() {
	for _, branch := range *b {
		branch.Delete()
	}
}

func (b *GitBranches) DefaultBranch() *GitBranch {
	for _, branch := range *b {
		if branch.IsDefault() {
			return &branch
		}
	}
	e.NewErr("failed to find the default branch")
	return nil
}

func (b *GitBranches) GetByName(branchName string) *GitBranch {
	for _, branch := range *b {
		if branch.Name == branchName {
			return &branch
		}
	}
	e.NewErr("failed as '%s' was not found", branchName)
	return nil
}
