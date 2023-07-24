package main

import (
	"self-encapsulated-logger/logger"
)

func main() {
	logger.SetFile("log/test.log")
	logger.SetLevel(logger.InfoLevel)
	logger.Info("hello world")
	logger.Debug("hello world -- Debug")
}
