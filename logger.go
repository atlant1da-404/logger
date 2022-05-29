package logger

import (
	"fmt"
	"log"
	"os"
)

type customLogger struct {
	CustomLogger
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
}

type CustomLogger interface {
	InfoLog(v ...interface{})
	WarnLog(v ...interface{})
	ErrorLog(v ...interface{})
	GetLogger() *customLogger
}

func NewCustomLogger() *customLogger {
	return &customLogger{
		info:  log.New(os.Stdout, "INFO: ", log.Ldate),
		warn:  log.New(os.Stdout, "WARN: ", log.Ldate),
		error: log.New(os.Stdout, "ERROR: ", log.Ldate),
	}
}

var instance *customLogger

var (
	red    = color("\033[1;31m%s\033[0m")
	yellow = color("\033[1;33m%s\033[0m")
	teal   = color("\033[1;36m%s\033[0m")
)

func (l *customLogger) InfoLog(v ...interface{}) {
	l.info.Println(teal(v...))
}

func (l *customLogger) WarnLog(v ...interface{}) {
	l.warn.Println(yellow(v...))
}

func (l *customLogger) ErrorLog(v ...interface{}) {
	l.error.Println(red(v...))
}

func (l *customLogger) GetLogger() *customLogger {
	return instance
}

func color(colorCode string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorCode,
			fmt.Sprint(args...))
	}
	return sprint
}
