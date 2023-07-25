package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	debugLogger   *log.Logger
	infoLogger    *log.Logger
	warnLogger    *log.Logger
	errorLogger   *log.Logger
	logLevel      int
	logging       *os.File
	day           int
	logFile       string
	dayChangeLock sync.RWMutex
	nowFunc       func() time.Time = time.Now
)

const (
	DebugLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

func SetNowFunc(f func() time.Time) { // for test
	nowFunc = f
}

func SetLogLevel(level int) {
	logLevel = level
}

func SetLogFile(file string) {
	logFile = file
	now := time.Now()
	var err error
	if logging, err = os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o664); err != nil {
		panic(err)
	} else {
		debugLogger = log.New(logging, "[DEBUG] ", log.LstdFlags)
		infoLogger = log.New(logging, "[INFO] ", log.LstdFlags)
		warnLogger = log.New(logging, "[WARN] ", log.LstdFlags)
		errorLogger = log.New(logging, "[ERROR] ", log.LstdFlags)
		day = now.YearDay()
		dayChangeLock = sync.RWMutex{}
	}
}

func checkAndChangeLogfile() {
	dayChangeLock.Lock()
	defer dayChangeLock.Unlock()
	now := nowFunc()
	if now.YearDay() == day {
		return
	}
	logging.Close()
	postFix := now.Add(-24 * time.Hour).Format("20060102")
	if err := os.Rename(logFile, logFile+"."+postFix); err != nil {
		fmt.Printf("append date postfix %s to log file %s failed: %v\n", postFix, logFile, err)
		return
	}
	var err error
	if logging, err = os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o664); err != nil {
		fmt.Printf("create log file %s failed %v\n", logFile, err)
		return
	} else {
		debugLogger = log.New(logging, "[DEBUG] ", log.LstdFlags)
		infoLogger = log.New(logging, "[INFO] ", log.LstdFlags)
		warnLogger = log.New(logging, "[WARN] ", log.LstdFlags)
		errorLogger = log.New(logging, "[ERROR] ", log.LstdFlags)
		day = now.YearDay()
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
		checkAndChangeLogfile()
		debugLogger.Printf(addPrefix()+" "+format, v...)
	}
}

func Info(format string, v ...any) {
	if logLevel <= InfoLevel {
		checkAndChangeLogfile()
		infoLogger.Printf(addPrefix()+" "+format, v...)
	}
}

func Warn(format string, v ...any) {
	if logLevel <= WarnLevel {
		checkAndChangeLogfile()
		warnLogger.Printf(addPrefix()+" "+format, v...)
	}
}

func Error(format string, v ...any) {
	if logLevel <= ErrorLevel {
		checkAndChangeLogfile()
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
