package logger

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type ServerFormatter struct {
	Timestamp string
	Level     logrus.Level
	Message   string
}

func (f *ServerFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format("2006-01-02T15:04:05.999-07:00")
	return []byte(fmt.Sprintf("%s %s %s\n", 
		timestamp, 
		entry.Level, 
		entry.Message, 
	)), nil
}