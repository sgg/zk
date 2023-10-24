package zk

import "log"

// Logger is an interface that can be implemented to provide custom log output.
type Logger interface {
	// Deprecated: use Debug
	Printf(string, ...interface{})
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

type defaultLogger struct{}

func (l defaultLogger) Debug(msg string, args ...interface{}) {
	log.Printf(msg, args...)
}

func (l defaultLogger) Info(msg string, args ...interface{}) {
	log.Printf(msg, args...)
}

func (l defaultLogger) Warn(msg string, args ...interface{}) {
	log.Printf(msg, args...)
}

func (l defaultLogger) Error(msg string, args ...interface{}) {
	log.Printf(msg, args...)
}

func (defaultLogger) Printf(msg string, args ...interface{}) {
	log.Printf(msg, args...)
}
