package config

import (
	"fmt"
	"path/filepath"

	e "github.com/delineateio/go-neat/errors"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type NeatConfig struct {
	viper *viper.Viper
}

var Config = GetConfig(NewGlobalConfig())

func GetConfig(info *ConfigInfo) NeatConfig {
	zerolog.SetGlobalLevel(zerolog.Disabled)
	a := NeatConfig{
		viper: viper.New(),
	}
	a.read(info)
	return a
}

func (a *NeatConfig) read(i *ConfigInfo) {

	v := a.viper

	for _, key := range i.defaults.Keys() {
		value, ok := i.defaults.Get(key)
		if ok {
			v.SetDefault(key, value)
			debugConfigSet(key, value)
		}
	}
	v.SetConfigName(i.filename)
	v.SetConfigType(i.extension)
	v.AddConfigPath(i.path)
	err := v.ReadInConfig()

	e.CheckIfError(err, "failed to add/merge '%s' config",
		filepath.Join(i.path, fmt.Sprintf("%s.%s",
			i.filename,
			i.extension)))
}

func (a *NeatConfig) GetString(key string) string {
	value := a.viper.GetString(key)
	debugConfigGet("config_get", "str", key, value)
	return value
}

func (a *NeatConfig) GetStrings(key string) []string {
	value := a.viper.GetStringSlice(key)
	debugConfigGet("config_get", "strs", key, value)
	return value
}

func (a *NeatConfig) GetInt(key string) int {
	value := a.viper.GetInt(key)
	debugConfigGet("config_get", "int", key, value)
	return value
}

func (a *NeatConfig) GetBool(key string) bool {
	value := a.viper.GetBool(key)
	debugConfigGet("config_get", "bool", key, value)
	return value
}
