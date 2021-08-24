package logger

import (
	"fmt"
	"log"
)

type Logger struct {
	loggerName string
}

func New(loggerName string) *Logger {
	return &Logger{
		loggerName: loggerName,
	}
}

func (l *Logger) Info(s interface{}) {
	log.Println(s)
}
func (l *Logger) Infof(s string, a ...interface{}) {
	log.Println(fmt.Sprintf(s, a...))
}
