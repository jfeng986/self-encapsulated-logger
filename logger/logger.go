package logger

import (
	"log"
	"os"
)

var infoLogger *log.Logger

var logFile *os.File

func SetFile(file string) {
	var err error
	logFile, err = os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o664)
	if err != nil {
		panic(err)
	} else {
		infoLogger = log.New(logFile, "[INFO]", log.LstdFlags|log.Lshortfile)
	}
}

func Info(format string, v ...any) {
	infoLogger.Printf(format, v...)
}
