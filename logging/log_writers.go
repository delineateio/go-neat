package logging

import (
	"io"
)

type logWriters map[string]io.Writer

func (w *logWriters) keys() []string {
	keys := make([]string, len(*w))
	i := 0
	for k := range *w {
		keys[i] = k
		i++
	}
	return keys
}

func (w *logWriters) values() []io.Writer {
	values := make([]io.Writer, len(*w))
	i := 0
	for _, v := range *w {
		values[i] = v
		i++
	}
	return values
}
