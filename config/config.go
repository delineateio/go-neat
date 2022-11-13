package config

import (
	"github.com/rs/zerolog"
)

func ReadGlobalConfig() {
	zerolog.SetGlobalLevel(zerolog.Disabled)
	config := NewGlobalConfig()
	config.WriteIfNotExists()
	config.Read()
}
