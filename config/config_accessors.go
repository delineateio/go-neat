package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

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
