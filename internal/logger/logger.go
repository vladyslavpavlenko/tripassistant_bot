package logger

import "fmt"

// Logger represents logger used to debug or error information
type Logger interface {
	Debugf(format string, args ...any)
	Errorf(format string, args ...any)
}

// MyLogger is a custom logger of type Logger
type MyLogger struct {
}

// Debugf implements the Debugf method of the Logger interface
func (l *MyLogger) Debugf(format string, args ...any) {
	fmt.Printf("[DEBUG] "+format+"\n", args...)
}

// Errorf implements the Errorf method of the Logger interface
func (l *MyLogger) Errorf(format string, args ...any) {
	fmt.Printf("[ERROR] "+format+"\n", args...)
}
