package adapter

import "fmt"

// ExternalLogger is a third-party logger you can't change.
type ExternalLogger struct{}

func (e *ExternalLogger) Info(msg string, fields map[string]string) {
	fmt.Println("[INFO] " + msg)
}

// Logger is the target interface used by the application.
type Logger interface {
	Info(message string)
}

// AppLogger is an internal logger.
type AppLogger struct{}

func (l *AppLogger) Info(msg string) {
	fmt.Printf("[App INFO] %s\n", msg)
}

// ThirdPartyLogger is an interface set to handle the external logger
type ThirdPartyLogger interface {
	Info(msg string, fields map[string]string)
}

// ThirdPartyAdapter adapts ExternalLogger to conform to Logger.
type ThirdPartyAdapter struct {
	logger ThirdPartyLogger
}

func (e *ThirdPartyAdapter) Info(message string) {
	e.logger.Info(message, map[string]string{
		"source": "app",
	})
}

func NewThirdPartyLoggerAdapter(logger ThirdPartyLogger) Logger {
	return &ThirdPartyAdapter{
		logger: logger,
	}
}

// Run demonstrates the Adapter pattern
func Run() {
	var logger Logger

	logger = &AppLogger{}
	logger.Info("App logger in action")

	logger = NewThirdPartyLoggerAdapter(&ExternalLogger{})
	logger.Info("Using third-party logger via adapter")
}
