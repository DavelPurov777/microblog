package logger

import "log"

type LogEvent struct {
	Level   string
	Message string
}

type Logger interface {
	Info(msg string)
	Error(msg string)
	Close()
}

type channelLogger struct {
	ch chan LogEvent
}

func NewLogger(buffer int) *channelLogger {
	l := &channelLogger{
		ch: make(chan LogEvent, buffer),
	}

	go l.listen()

	return l
}

func (l *channelLogger) listen() {
	for event := range l.ch {
		log.Printf("[%s] %s", event.Level, event.Message)
	}
}

func (l *channelLogger) Info(msg string) {
	l.ch <- LogEvent{"INFO", msg}
}

func (l *channelLogger) Error(msg string) {
	l.ch <- LogEvent{"ERROR", msg}
}

func (l *channelLogger) Close() {
	close(l.ch)
}
