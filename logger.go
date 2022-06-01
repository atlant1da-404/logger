/*

How to use this package:


1. Call a new instance of customLogger:           lg := logger.NewCustomLogger()
2. Instance return interface{} with several public methods:
	ex: lg.Info("some text") -> print in console a data with teal color  -> output like "INFO: some text"
	ex: lg.Warn("some text", error) -> print in console a data with yellow color -> output like "WARN: some text [error]"
3. Write in file:
    ex: lg.File("filename.log", "some text") -> create or append to file a "some text"
4: Write in file 2 way:
    ex: lg.Info("some text").File("filename.log") -> also work, file created or append to file a "some text" and prints in console
5: Custom Colors:
	ex: lg.Prefix(logger.Black, "Some Tag(Prefix): ").Console("some text") -> output like "Some Tag(Prefix): some text"
6: Write in file by custom color:
    ex: lg.Prefix(logger.Black, "Some Tag(Prefix): ").Console("hh").File("hello.log") -> custom color in console and also created or append to file info
7: FileF():
	ex: lg.FileF("filename.log", "[some text]", err.Error()) -> format wrapper for error, output -> [some text] [ err ]

Todo ex:

	lg := logger.NewCustomLogger()

	lg.Info("Information console block!")
>[INFO] Information console block!

	lg.Warn("Something wrong")
>[WARN] Something wrong

	lg.Error("Crash!!!!")
>[ERROR] Crash!!!!

	lg.Info("need to store in logs").File("store.log")
>Create or Append to file:
	[INFO] need to store in logs

	lg.Info("[Code 2503], err.Error())
>[INFO] [Code 2503] [err]

	lg.File("store.log", "some data")
>Create or Append to file:
	[some data]

	lg.FileF("store.log", "[Code 2503]", err.Error())
>Create or Append to file a format error:
	[Code 2503] [err]
*/

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
		Info(msg string, v ...interface{}) *customLogger
		Warn(msg string, v ...interface{}) *customLogger
		Error(msg string, v ...interface{}) *customLogger
		Debug(msg string, v ...interface{}) *customLogger
		Console(v ...interface{}) *customLogger
		Prefix(color func(...interface{}) string, prefix string) *customLogger
		File(fileName string, v ...interface{})
		FileF(fileName string, msg string, v ...interface{})
	}
)

// fileOs is instance of file
var fileOs *os.File

// NewCustomLogger initializing a new instance of logger
func NewCustomLogger() *customLogger {
	return &customLogger{
		info:  log.New(os.Stdout, Teal("[INFO] "), log.Ldate),
		warn:  log.New(os.Stdout, Yellow("[WARN] "), log.Ldate),
		error: log.New(os.Stdout, Red("[ERROR] "), log.Ldate),
		debug: log.New(os.Stdout, Green("[DEBUG] "), log.Ldate),
	}
}

// Info Print in console the data with prefix INFO: and teal color
func (l *customLogger) Info(msg string, v ...interface{}) *customLogger {
	l.info.Println(msg, v)
	l.tempData = fmt.Sprintf("%v", v)
	return l
}

// Warn Print in console the data with prefix WARN: and yellow color
func (l *customLogger) Warn(msg string, v ...interface{}) *customLogger {
	l.warn.Println(msg, v)
	l.tempData = fmt.Sprintf("%v", v)
	return l
}

// Error Print in console the data with prefix ERROR: and red color
func (l *customLogger) Error(msg string, v ...interface{}) *customLogger {
	l.error.Println(msg, v)
	l.tempData = fmt.Sprintf("%v", v)
	return l
}

// Debug Print in console the data with prefix DEBUG: and green color
func (l *customLogger) Debug(msg string, v ...interface{}) *customLogger {
	l.debug.Println(msg, v)
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
func (l *customLogger) File(fileName string, v ...interface{}) {
	if len(v) != 0 {
		l.tempData = fmt.Sprintf("%v", v)
	}

	// if we are logging in another file we are out from previous (close)
	if fileOs != nil && fileOs.Name() != fileName {
		_ = fileOs.Close()
	}

	f, _ := os.OpenFile(fileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	fileOs = f
	logger := log.New(f, "", log.LstdFlags)
	logger.Println(l.tempData)
}

// FileF write a data in file with message
func (l *customLogger) FileF(fileName, msg string, v ...interface{}) {
	if len(v) != 0 {
		l.tempData = fmt.Sprintf("%v", v)
	}

	// if we are logging in another file we are out from previous (close)
	if fileOs != nil && fileOs.Name() != fileName {
		_ = fileOs.Close()
	}

	f, _ := os.OpenFile(fileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	fileOs = f
	logger := log.New(f, "", log.LstdFlags)
	logger.Println(msg, l.tempData)
}
