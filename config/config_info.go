package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	e "github.com/delineateio/go-neat/errors"
	o "github.com/elliotchance/orderedmap/v2"
	"github.com/spf13/viper"
)

type ConfigInfo struct {
	filename  string
	path      string
	extension string
	defaults  *o.OrderedMap[string, any]
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

func (i *ConfigInfo) Merge() *ConfigInfo {
	addConfig(i, viper.MergeInConfig)
	return i
}

func (i *ConfigInfo) Read() *ConfigInfo {
	addConfig(i, viper.ReadInConfig)
	return i
}

func addConfig(info *ConfigInfo, method func() error) {
	for _, key := range info.defaults.Keys() {
		value, ok := info.defaults.Get(key)
		if ok {
			viper.SetDefault(key, value)
			debugConfigSet(key, value)
		}
	}
	viper.SetConfigName(info.filename)
	viper.SetConfigType(info.extension)
	viper.AddConfigPath(info.path)
	err := method()
	e.CheckIfError(err, "failed to add '%s' config",
		filepath.Join(info.path, fmt.Sprintf("%s.%s",
			info.filename,
			info.extension)))
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
