package logger

import (
	"log"
	"os"
)

// Logger represents a logger with different log levels
type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

// New creates a new logger instance
func New() *Logger {
	return &Logger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Info logs informational messages
func (l *Logger) Info(v ...interface{}) {
	l.infoLogger.Println(v...)
}

// Error logs error messages
func (l *Logger) Error(v ...interface{}) {
	l.errorLogger.Println(v...)
}

// Fatal logs fatal error messages and exits the application
func (l *Logger) Fatal(v ...interface{}) {
	l.errorLogger.Fatal(v...)
}

// Global logger instance
var loggerInstance = New()

// Exported functions for global logging
func Info(v ...interface{}) {
	loggerInstance.Info(v...)
}

func Error(v ...interface{}) {
	loggerInstance.Error(v...)
}

func Fatal(v ...interface{}) {
	loggerInstance.Fatal(v...)
}
