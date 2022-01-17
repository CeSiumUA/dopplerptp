package logging

import "fmt"

type Logger struct {
}

var GlobalLogger Logger = Logger{}

func (lg *Logger) LogInfo(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	fmt.Print("\n")
}
