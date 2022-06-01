package logger

import (
	"fmt"
	"log"
	"os"
)

type customLogger struct {
	CustomLogger
	info     *log.Logger
	warn     *log.Logger
	error    *log.Logger
	debug    *log.Logger
	tempData string
	prefix   string
	color    func(...interface{}) string
}

type CustomLogger interface {
	Info(v ...interface{}) *customLogger
	Warn(v ...interface{}) *customLogger
	Error(v ...interface{}) *customLogger
	Debug(v ...interface{}) *customLogger
	File(fileName string, v ...interface{})
	Prefix(color func(...interface{}) string, prefix string) *customLogger
	Console(v ...interface{}) *customLogger
}

var fileOs *os.File

func NewCustomLogger() *customLogger {
	return &customLogger{
		info:  log.New(os.Stdout, Teal("INFO: "), log.Ldate),
		warn:  log.New(os.Stdout, Yellow("WARN: "), log.Ldate),
		error: log.New(os.Stdout, Red("ERROR: "), log.Ldate),
		debug: log.New(os.Stdout, Green("DEBUG: "), log.Ldate),
	}
}

func (l *customLogger) Info(v ...interface{}) *customLogger {
	l.info.Println(v...)
	l.tempData = fmt.Sprintf("%v", v)
	return l
}

func (l *customLogger) Warn(v ...interface{}) *customLogger {
	l.warn.Println(v...)
	l.tempData = fmt.Sprintf("%v", v)
	return l
}

func (l *customLogger) Error(v ...interface{}) *customLogger {
	l.error.Println(v...)
	l.tempData = fmt.Sprintf("%v", v)
	return l
}

func (l *customLogger) Debug(v ...interface{}) *customLogger {
	l.debug.Println(v...)
	l.tempData = fmt.Sprintf("%v", v)
	return l
}

func (l *customLogger) Prefix(color func(...interface{}) string, prefix string) *customLogger {
	l.prefix = prefix
	l.color = color
	return l
}

func (l *customLogger) Console(v ...interface{}) *customLogger {
	l.tempData = fmt.Sprintf("%v", v)

	if len([]rune(l.prefix)) == 0 || l.color == nil {
		l.color = White
		l.prefix = "DEFAULT: "
	}

	log.New(os.Stdout, l.color(l.prefix), log.Ldate).Println(v...)

	return l
}

func (l *customLogger) File(file string, v ...interface{}) {
	if len(v) != 0 {
		l.tempData = fmt.Sprintf("%v", v)
	}

	// if we are logging in another file we are out from previous (close)
	if fileOs != nil && fileOs.Name() != file {
		fileOs.Close()
	}

	f, _ := os.OpenFile(file,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	fileOs = f
	logger := log.New(f, "", log.LstdFlags)
	logger.Println(l.tempData)
}
