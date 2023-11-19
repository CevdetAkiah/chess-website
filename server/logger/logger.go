package custom_log

import (
	"log"
	"os"
)

type Logger struct {
	log *log.Logger
}

type MagicLogger interface {
	Info(v string)
	Infof(format string, a ...any)
	Error(v string)
}

func NewLogger() *Logger {
	return &Logger{log: log.New(os.Stdout, "database-logger ", log.Flags())}
}

func (l *Logger) Info(v string) {
	// a normal print
	l.log.Println(v)
}

func (l *Logger) Infof(format string, a ...any) {
	l.log.Printf(format, a...)
}

func (l *Logger) Error(v string) {
	l.log.Println("ERROR: ", v)

}
