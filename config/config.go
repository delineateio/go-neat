package config

import (
	"fmt"
	"path/filepath"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	e "go.delineate.io/neat/errors"
)

func ReadGlobalConfig() {
	zerolog.SetGlobalLevel(zerolog.Disabled)
	config := NewGlobalConfigInfo()
	config.WriteIfNotExists()
	addConfig(config,
		viper.ReadInConfig,
	)
}

func MergeConfig(info *ConfigInfo) {
	addConfig(info,
		viper.MergeInConfig)
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
		filepath.Join(info.path, fmt.Sprintf("%s.%s", info.filename, info.extension)))
}

func GetString(key string) string {
	value := viper.GetString(key)
	debugConfigGet("config_get", "str", key, value)
	return value
}

func GetStrings(key string) []string {
	value := viper.GetStringSlice(key)
	debugConfigGet("config_get", "strs", key, value)
	return value
}

func GetInt(key string) int {
	value := viper.GetInt(key)
	debugConfigGet("config_get", "int", key, value)
	return value
}

func GetBool(key string) bool {
	value := viper.GetBool(key)
	debugConfigGet("config_get", "bool", key, value)
	return value
}

func debugConfigSet(key string, value any) {
	log.Debug().
		Str("event", "config_set").
		Str("key", key).
		Interface("value", value).
		Send()
}

func debugConfigGet(event, config_type, key string, value any) {
	log.Debug().
		Str("event", event).
		Str("type", config_type).
		Str("key", key).
		Interface("value", value).
		Send()
}
