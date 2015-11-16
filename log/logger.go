package log

import (
	"fmt"
	"io"
	"os"
)

type Logger interface {
	Info(msg string)
	Debug(msg string)
}

type DefaultLogger struct {
	out   io.Writer
	debug bool
}

func NewLogger(debug bool) Logger {
	return &DefaultLogger{
		out:   os.Stdout,
		debug: debug,
	}
}

func (lgr *DefaultLogger) Info(msg string) {
	fmt.Fprintln(lgr.out, msg)
}

func (lgr *DefaultLogger) Debug(msg string) {
	if lgr.debug {
		fmt.Fprintln(lgr.out, msg)
	}
}
