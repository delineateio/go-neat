package config

import (
	"errors"
	"fmt"
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

func NewDefaultConfigInfo(filename, path string) *ConfigInfo {
	return &ConfigInfo{
		filename:  filename,
		path:      path,
		extension: "yaml",
		defaults:  o.NewOrderedMap[string, any](),
	}
}

func NewGlobalConfigInfo() *ConfigInfo {
	dirname, err := os.UserHomeDir()
	e.CheckIfError(err, "failed to retrieve the home directory for the current user")

	return NewDefaultConfigInfo(".neat", filepath.Join(dirname, ".config")).
		AddDefault("log.level", "info").
		AddDefault("log.sinks", []string{"file"}).
		AddDefault("log.file.dir", "./.neat/logs").
		AddDefault("log.file.size_mb", 1).
		AddDefault("log.file.backups", 3).
		AddDefault("log.file.age_days", 7)
}

func (c *ConfigInfo) AddDefault(key string, value any) *ConfigInfo {
	c.defaults.Set(key, value)
	return c
}

func (c *ConfigInfo) WriteIfNotExists() {

	filename := fmt.Sprintf("%s.%s", c.filename, c.extension)
	configFile := filepath.Join(c.path, filename)

	if _, err := os.Stat(configFile); errors.Is(err, os.ErrNotExist) {
		fmt.Println(configFile)
	}
}
