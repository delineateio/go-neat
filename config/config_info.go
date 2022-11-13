package config

import (
	"os"
	"path/filepath"

	e "github.com/delineateio/go-neat/errors"
	o "github.com/elliotchance/orderedmap/v2"
)

type ConfigInfo struct {
	filename  string
	path      string
	extension string
	defaults  *o.OrderedMap[string, any]
}

func (c *ConfigInfo) AddDefault(key string, value any) *ConfigInfo {
	c.defaults.Set(key, value)
	return c
}

func NewDefaultConfig(filename, path string) *ConfigInfo {
	return &ConfigInfo{
		filename:  filename,
		path:      path,
		extension: "yaml",
		defaults:  o.NewOrderedMap[string, any](),
	}
}

func NewGlobalConfig() *ConfigInfo {
	dirname, err := os.UserHomeDir()
	e.CheckIfError(err, "failed to retrieve the home directory for the current user")

	return NewDefaultConfig(".neat", filepath.Join(dirname, ".config")).
		AddDefault("log.level", "info").
		AddDefault("log.sinks", []string{"file"}).
		AddDefault("log.file.dir", "./.neat/logs").
		AddDefault("log.file.size_mb", 1).
		AddDefault("log.file.backups", 3).
		AddDefault("log.file.age_days", 7)
}
