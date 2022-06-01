package logger

import (
	"fmt"
	"log"
	"os"
)

type (
	// customLogger is private struct with private fields and public methods
	customLogger struct {
		CustomLogger
		info     *log.Logger
		warn     *log.Logger
		error    *log.Logger
		debug    *log.Logger
		tempData string
		prefix   string
		color    func(...interface{}) string
	}
	// CustomLogger - interface of customLogger private struct with public methods
	CustomLogger interface {
		Info(v ...interface{}) *customLogger
		Warn(v ...interface{}) *customLogger
		Error(v ...interface{}) *customLogger
		Debug(v ...interface{}) *customLogger
		Console(v ...interface{}) *customLogger
		Prefix(color func(...interface{}) string, prefix string) *customLogger
		File(fileName string, v ...interface{})
	}
)

// fileOs is instance of file
var fileOs *os.File

// NewCustomLogger initializing a new instance of logger
func NewCustomLogger() *customLogger {
	return &customLogger{
		info:  log.New(os.Stdout, Teal("INFO: "), log.Ldate),
		warn:  log.New(os.Stdout, Yellow("WARN: "), log.Ldate),
		error: log.New(os.Stdout, Red("ERROR: "), log.Ldate),
		debug: log.New(os.Stdout, Green("DEBUG: "), log.Ldate),
	}
}

// Info Print in console the data with prefix INFO: and teal color
func (l *customLogger) Info(v ...interface{}) *customLogger {
	l.info.Println(v...)
	l.tempData = fmt.Sprintf("%v", v)
	return l
}

// Warn Print in console the data with prefix WARN: and yellow color
func (l *customLogger) Warn(v ...interface{}) *customLogger {
	l.warn.Println(v...)
	l.tempData = fmt.Sprintf("%v", v)
	return l
}

// Error Print in console the data with prefix ERROR: and red color
func (l *customLogger) Error(v ...interface{}) *customLogger {
	l.error.Println(v...)
	l.tempData = fmt.Sprintf("%v", v)
	return l
}

// Debug Print in console the data with prefix DEBUG: and green color
func (l *customLogger) Debug(v ...interface{}) *customLogger {
	l.debug.Println(v...)
	l.tempData = fmt.Sprintf("%v", v)
	return l
}

// Prefix Custom method to make your own custom color with prefix
func (l *customLogger) Prefix(color func(...interface{}) string, prefix string) *customLogger {
	l.prefix = prefix
	l.color = color
	return l
}

// Console method print the data with custom prefix and color
func (l *customLogger) Console(v ...interface{}) *customLogger {
	l.tempData = fmt.Sprintf("%v", v)

	if len([]rune(l.prefix)) == 0 || l.color == nil {
		l.color = White
		l.prefix = "DEFAULT: "
	}

	log.New(os.Stdout, l.color(l.prefix), log.Ldate).Println(v...)
	return l
}

// File write data in file
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
