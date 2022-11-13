package config

import "github.com/rs/zerolog/log"

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
