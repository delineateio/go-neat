package config

import (
	"embed"
	"errors"
	"os"
	"path/filepath"

	e "github.com/delineateio/go-neat/errors"
	u "github.com/delineateio/go-neat/ui"
)

//go:embed .neat-global.yaml
var gcf embed.FS

//go:embed .neat-repo.yaml
var rcf embed.FS

func initConfig(resourceName string, source embed.FS, basepath string, parts ...string) {

	parts = append(parts, ".neat.yaml")
	path := filepath.Join(parts...)

	path, err := filepath.Abs(path)
	e.CheckIfError(err, "failed to get the absolute path for '%s'", path)

	relPath, err := filepath.Rel(basepath, path)
	e.CheckIfError(err, "failed to get the relative path for '%s'", path)

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		data, err := source.ReadFile(resourceName)
		e.CheckIfError(err, "failed to read resource '%s'", resourceName)

		f, err := os.Create(path)
		e.CheckIfError(err, "failed to create file '%s'", relPath)
		defer f.Close()

		_, err = f.Write(data)
		e.CheckIfError(err, "failed to write to '%s'", relPath)

		err = f.Sync()
		e.CheckIfError(err, "failed to write to '%s'", relPath)

		u.Successful("successfully created '%s'", relPath)

	} else {
		u.Skipped("skipped as '%s' exists", relPath)
	}
}

func WriteGlobalInitConfig() {

	home, err := os.UserHomeDir()
	e.CheckIfError(err, "failed to get the current user home directory")

	initConfig(".neat-global.yaml",
		gcf,
		home,
		home, ".config")
}

func WriteRepoInitConfig(path string) {

	repoPath, err := filepath.Abs(path)
	e.CheckIfError(err, "failed to get the absolute path for '%s'", path)

	initConfig(".neat-repo.yaml",
		rcf,
		repoPath,
		repoPath)
}
