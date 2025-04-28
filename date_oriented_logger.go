package logger

import (
	"os"
	"path/filepath"
	"time"
)

type DateOrientedLogger struct {
	Logger
	logFile     *os.File
	logFileName string
}

func (l *DateOrientedLogger) _setup(curTime time.Time) error {
	filename := curTime.Format(time.DateOnly) + ".log"

	if l.logFileName == filename && l.logFile != nil {
		return nil
	}

	l.logFileName = filename

	logFile, err := os.OpenFile(filepath.Join(l.fullpath, l.logFileName), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	l.logFile = logFile

	return nil
}

func (l *DateOrientedLogger) setup() {
	if err := l._setup(time.Now()); err != nil {
		printAndExit(err.Error(), 1)
	}
}

func (l *DateOrientedLogger) _Log(message string, level LogLevel, curTime time.Time) error {
	if l.level > level {
		return nil
	}

	if err := l._setup(curTime); err != nil {
		return err
	}

	output := ""

	for _, token := range *l.formatter.orderedTokens {
		val, err := getFormatterTokenValue(token, message, level)
		if err != nil {
			return err
		}
		output += val
	}

	if _, err := l.logFile.Write([]byte(output + "\n")); err != nil {
		return err
	}

	return nil
}

func (l *DateOrientedLogger) Log(message string, level LogLevel) {
	if err := l._Log(message, level, time.Now()); err != nil {
		printAndExit(err.Error(), 1)
	}
}

func (l *DateOrientedLogger) DEBUG(message string) {
	l.Log(message, DEBUG)
}

func (l *DateOrientedLogger) INFO(message string) {
	l.Log(message, INFO)
}

func (l *DateOrientedLogger) WARN(message string) {
	l.Log(message, WARN)
}

func (l *DateOrientedLogger) ERROR(message string) {
	l.Log(message, ERROR)
}
