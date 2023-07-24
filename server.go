package main

import (
	"self-encapsulated-logger/logger"
)

func main() {
	// logger.SetLogFile()
	logger.SetLogLevel(logger.DebugLevel)
	logger.Info("hello world -- Info")
	logger.Debug("hello world -- Debug")
	logger.Warn("hello world -- Warn")
	logger.Error("hello world -- Error")
}
