package logger

import "fmt"

// Colors
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

func color(colorCode string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorCode,
			fmt.Sprint(args...))
	}
	return sprint
}
