package main

import (
	"self-encapsulated-logger/logger"
)

func main() {
	logger.SetFile("log/test.log")
	logger.Info("hello world")
}
