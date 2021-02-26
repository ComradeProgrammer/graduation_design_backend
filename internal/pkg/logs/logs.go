package logs

import (
	"fmt"
	"runtime"
	"time"
)

const (
	red    int = 31
	green  int = 32
	yellow int = 33
	blue   int = 34
	aqua   int = 36
	white  int = 37
)

func Info(format string, objs ...interface{}) {
	str := fmt.Sprintf(format, objs...)
	_, filename, line, ok := runtime.Caller(1)
	time := time.Now().Format("2006-01-02 15:04:05.0000")
	var logString string
	if ok {
		logString = fmt.Sprintf("[Info   ]%s %s:%d %s\n", time, filename, line, str)
	} else {
		logString = fmt.Sprintf("[Info   ]%s unknown location %s\n", time, str)
	}
	fmt.Print(colorString(logString, blue))

}
func Warning(format string, objs ...interface{}) {
	str := fmt.Sprintf(format, objs...)
	_, filename, line, ok := runtime.Caller(1)
	time := time.Now().Format("2006-01-02 15:04:05.0000")
	var logString string
	if ok {
		logString = fmt.Sprintf("[Warning]%s %s:%d %s\n", time, filename, line, str)
	} else {
		logString = fmt.Sprintf("[Warning]%s unknown location %s\n", time, str)
	}
	fmt.Print(colorString(logString, yellow))

}
func Error(format string, objs ...interface{}) {
	str := fmt.Sprintf(format, objs...)
	_, filename, line, ok := runtime.Caller(1)
	time := time.Now().Format("2006-01-02 15:04:05.0000")
	var logString string
	if ok {
		logString = fmt.Sprintf("[Error  ]%s %s:%d %s\n", time, filename, line, str)
	} else {
		logString = fmt.Sprintf("[Error  ]%s unknown location %s\n", time, str)
	}
	fmt.Print(colorString(logString, red))

}
func Fatal(format string, objs ...interface{}) {
	str := fmt.Sprintf(format, objs...)
	_, filename, line, ok := runtime.Caller(1)
	time := time.Now().Format("2006-01-02 15:04:05.0000")
	var logString string
	if ok {
		logString = fmt.Sprintf("[Fatal  ]%s %s:%d %s\n", time, filename, line, str)
	} else {
		logString = fmt.Sprintf("[Fatal  ]%s unknown location %s\n", time, str)
	}
	fmt.Print(colorString(logString, aqua))

}

func colorString(str string, color int) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, str)
}
