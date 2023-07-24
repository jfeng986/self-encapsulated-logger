package logger

import (
	"log"
	"os"
)

var (
	infoLogger  *log.Logger
	debugLogger *log.Logger
	logLevel    int
)

const (
	DebugLevel = iota
	InfoLevel
)

func SetLevel(level int) {
	logLevel = level
}

func SetFile(file string) {
	logFile, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o664)
	if err != nil {
		panic(err)
	} else {
		infoLogger = log.New(logFile, "[INFO] ", log.LstdFlags|log.Lshortfile)
		debugLogger = log.New(logFile, "[DEBUG] ", log.LstdFlags|log.Lshortfile)
	}
}

func Info(format string, v ...any) {
	if logLevel <= InfoLevel {
		infoLogger.Printf(format, v...)
	}
}

func Debug(format string, v ...any) {
	if logLevel <= DebugLevel {
		debugLogger.Printf(format, v...)
	}
}
