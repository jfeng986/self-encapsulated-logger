package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	logLevel    int
	logFile     string
)

const (
	DebugLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

func SetLogLevel(level int) {
	logLevel = level
}

func init() {
	now := time.Now()
	formattedDate := now.Format("2006-01-02")
	logFile = fmt.Sprintf("log%s%s.log", string(os.PathSeparator), formattedDate)
	logging, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o664)
	if err != nil {
		panic(err)
	} else {
		debugLogger = log.New(logging, "[DEBUG] ", log.LstdFlags)
		infoLogger = log.New(logging, "[INFO] ", log.LstdFlags)
		warnLogger = log.New(logging, "[WARN] ", log.LstdFlags)
		errorLogger = log.New(logging, "[ERROR] ", log.LstdFlags)
	}
}

func addPrefix() string {
	file, _, line := getLineNo()
	arr := strings.Split(file, "/")
	if len(arr) > 3 {
		arr = arr[len(arr)-3:]
	}
	return strings.Join(arr, "/") + ":" + strconv.Itoa(line)
}

func Debug(format string, v ...any) {
	if logLevel <= DebugLevel {
		debugLogger.Printf(addPrefix()+" "+format, v...)
	}
}

func Info(format string, v ...any) {
	if logLevel <= InfoLevel {
		infoLogger.Printf(addPrefix()+" "+format, v...)
	}
}

func Warn(format string, v ...any) {
	if logLevel <= WarnLevel {
		warnLogger.Printf(addPrefix()+" "+format, v...)
	}
}

func Error(format string, v ...any) {
	if logLevel <= ErrorLevel {
		errorLogger.Printf(addPrefix()+" "+format, v...)
	}
}

func getLineNo() (string, string, int) {
	funcName, file, line, ok := runtime.Caller(3)
	if ok {
		return file, runtime.FuncForPC(funcName).Name(), line
	} else {
		return "", "", 0
	}
}
