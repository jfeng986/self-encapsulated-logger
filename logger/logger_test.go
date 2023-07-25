package logger

import (
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	SetLogLevel(DebugLevel)

	SetLogFile("../log/test.log")

	Debug("This is a debug log for today")
	Info("This is an info log for today")
	Warn("This is a warning log for today")
	Error("This is an error log for today")

	SetNowFunc(func() time.Time {
		return time.Now().Add(24 * time.Hour)
	})

	Debug("This is a debug log for tomorrow")
	Info("This is an info log for tomorrow")
	Warn("This is a warning log for tomorrow")
	Error("This is an error log for tomorrow")
}
