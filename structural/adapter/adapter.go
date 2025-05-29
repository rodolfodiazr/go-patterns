package adapter

import "fmt"

// Logger is the target interface used by the application.
type Logger interface {
	Info(message string)
}

// AppLogger is an internal logger.
type AppLogger struct{}

func (l *AppLogger) Info(msg string) {
	fmt.Printf("[App INFO] %s\n", msg)
}

// ExternalLogger is a third-party logger you can't change.
type ExternalLogger struct{}

func (e *ExternalLogger) Info(msg string, fields map[string]string) {
	fmt.Println("[ExternalLogger]", msg)
}

// ExternalLoggerAdapter adapts ExternalLogger to conform to Logger.
type ExternalLoggerAdapter struct {
	logger *ExternalLogger
}

func NewExternalLoggerAdapter(logger *ExternalLogger) Logger {
	return &ExternalLoggerAdapter{
		logger: logger,
	}
}

func (e *ExternalLoggerAdapter) Info(message string) {
	e.logger.Info(message, map[string]string{
		"source": "app",
	})
}

// Run demonstrates the Adapter pattern
func Run() {
	var logger Logger

	logger = &AppLogger{}
	logger.Info("App logger in action")

	logger = NewExternalLoggerAdapter(&ExternalLogger{})
	logger.Info("Using third-party logger via adapter")
}
