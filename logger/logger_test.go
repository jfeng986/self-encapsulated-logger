package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	SetLogLevel(DebugLevel)
	// SetLogFile("../log/test.log")

	Debug("This is a debug message: %s", "debug")
	Info("This is an info message: %s", "info")
	Warn("This is a warning message: %s", "warning")
	Error("This is an error message: %s", "error")

	Debug("This is a debug message for the next day: %s", "debug")
	Info("This is an info message for the next day: %s", "info")
	Warn("This is a warning message for the next day: %s", "warning")
	Error("This is an error message for the next day: %s", "error")
}
