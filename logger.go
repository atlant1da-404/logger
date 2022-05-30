package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	Black  = color("\033[1;30m%s\033[0m")
	Red    = color("\033[1;31m%s\033[0m")
	Green  = color("\033[1;32m%s\033[0m")
	Yellow = color("\033[1;33m%s\033[0m")
	Blue   = color("\033[1;34m%s\033[0m")
	Purple = color("\033[1;35m%s\033[0m")
	Teal   = color("\033[1;36m%s\033[0m")
	White  = color("\033[1;37m%s\033[0m")
)

type customLogger struct {
	CustomLogger
	info     *log.Logger
	warn     *log.Logger
	error    *log.Logger
	debug    *log.Logger
	tempData string
}

type customColors struct {
	CustomColors
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
}

type CustomColors interface {
	Prefix(color func(...interface{}) string, prefix string) *customColors
	Console(v ...interface{}) *customLogger
}

var (
	loggerInstance *customLogger
	colorsInstance *customColors
	fileOs         *os.File
)

func NewCustomLogger() *customLogger {

	loggerInstance = &customLogger{
		info:  log.New(os.Stdout, Teal("INFO: "), log.Ldate),
		warn:  log.New(os.Stdout, Yellow("WARN: "), log.Ldate),
		error: log.New(os.Stdout, Red("ERROR: "), log.Ldate),
		debug: log.New(os.Stdout, Green("DEBUG: "), log.Ldate),
	}

	return loggerInstance
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

func NewCustomColorsLogger() *customColors {
	colorsInstance = &customColors{}
	return colorsInstance
}

func (c *customColors) Prefix(color func(...interface{}) string, prefix string) *customColors {
	c.prefix = prefix
	c.color = color
	return c
}

func (c *customColors) Console(v ...interface{}) *customLogger {
	logger := NewCustomLogger()
	logger.tempData = fmt.Sprintf("%v", v)
	log.New(os.Stdout, c.color(c.prefix), log.Ldate).Println(v...)

	return logger
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

func color(colorCode string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorCode,
			fmt.Sprint(args...))
	}
	return sprint
}
