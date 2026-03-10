package logger

import (
	"log"
)

type Logger interface {
	Info(msg string)
	Error(msg string)
}

type StdLogger struct{}

func NewLogger() *StdLogger {
	return &StdLogger{}
}

func (l *StdLogger) Info(msg string) {
	log.Printf("[INFO] %s", msg)
}

func (l *StdLogger) Error(msg string) {
	log.Printf("[ERROR] %s", msg)
}
