package logging

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	c "github.com/delineateio/go-neat/config"
	e "github.com/delineateio/go-neat/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

type logConfig struct {
	sinks      []string
	directory  string
	filename   string
	maxSize    int
	maxBackups int
	maxAge     int
	writers    logWriters
}

func (l *logConfig) hasSink(key string) bool {
	for _, k := range l.sinks {
		if strings.EqualFold(k, key) {
			return true
		}
	}
	return false
}

func NewLogConfig() *logConfig {
	return &logConfig{
		sinks:      c.GetStrings("log.sinks"),
		directory:  c.GetString("log.file.dir"),
		filename:   c.GetString("log.file.filename"),
		maxSize:    c.GetInt("log.file.size_mb"),
		maxBackups: c.GetInt("log.file.backups"),
		maxAge:     c.GetInt("log.file.age_days"),
		writers:    make(map[string]io.Writer, 0),
	}
}

func (l *logConfig) addConsole() *logConfig {
	if l.hasSink("console") {
		l.writers["console"] = zerolog.ConsoleWriter{Out: os.Stderr}
	}
	return l
}

func (l *logConfig) addFile() *logConfig {

	if l.hasSink("file") {
		home, err := os.UserHomeDir()
		e.CheckIfError(err, "failed to get the current users home directory")
		path := filepath.Join(home, l.directory, l.filename)
		l.writers["file"] = &lumberjack.Logger{
			Filename:   path,
			MaxBackups: l.maxBackups,
			MaxSize:    l.maxSize,
			MaxAge:     l.maxAge,
		}
	}
	return l
}

func (l *logConfig) configure() *logConfig {

	value := c.GetString("log.level")
	level, err := zerolog.ParseLevel(value)
	e.CheckIfError(err, "failed to parse '%s' to log level", value)

	zerolog.SetGlobalLevel(level)

	mw := io.MultiWriter(l.writers.values()...)
	log.Logger = zerolog.New(mw).With().Timestamp().Logger()
	return l
}

func (l *logConfig) log() *logConfig {
	event := log.Info().
		Str("event", "logger_initialised").
		Strs("writers", l.writers.keys())
	if l.writers["file"] != nil {
		event.Str("dir", l.directory).
			Int("size_mb", l.maxSize).
			Int("backups", l.maxBackups).
			Int("age_days", l.maxAge)
	}
	event.Send()
	return l
}
