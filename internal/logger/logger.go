package logger

import "log"

type LogEvent struct {
	Level   string
	Message string
}

type Logger struct {
	ch chan LogEvent
}

func NewLogger(buffer int) *Logger {
	l := &Logger{
		ch: make(chan LogEvent, buffer),
	}

	go l.listen()

	return l
}

func (l *Logger) listen() {
	for event := range l.ch {
		log.Printf("[%s] %s", event.Level, event.Message)
	}
}

func (l *Logger) Info(msg string) {
	l.ch <- LogEvent{"INFO", msg}
}

func (l *Logger) Error(msg string) {
	l.ch <- LogEvent{"ERROR", msg}
}

func (l *Logger) Close() {
	close(l.ch)
}
