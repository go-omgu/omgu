package log

import (
	"fmt"
	"io"
	"strings"

	"github.com/fsamin/go-dump"
)

type Logger struct {
	w         io.Writer
	delimiter string
	level     int
}

const (
	Normal = iota
	Verbose
	Debug
)

func New(w io.Writer, level int) *Logger {
	l := &Logger{w: w, level: level}
	l.delimiter = " :: "
	return l
}

func (log *Logger) SetDelim(d string) {
	log.delimiter = d
}

func (log *Logger) output(level int, str ...string) error {
	if level > log.level {
		return nil
	}
	msg := strings.Join(str, log.delimiter)
	_, err := fmt.Fprintf(log.w, "%s\n", msg)
	return err
}

func (log *Logger) Log(str ...string) error {
	return log.output(Normal, str...)
}

func (log *Logger) Verbose(str ...string) error {
	return log.output(Verbose, str...)
}

func (log *Logger) Debug(str ...string) error {
	return log.output(Debug, str...)
}

func (log *Logger) Dump(v interface{}) error {
	m, err := dump.ToStringMap(v)
	if err != nil {
		return err
	}

	for k, v := range m {
		err = log.output(Debug, k, v)
		if err != nil {
			return err
		}
	}

	return nil
}
