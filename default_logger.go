package logger

import (
	"os"
	"path/filepath"
)

type DefaultLogger struct {
	Logger
	logFile *os.File
}

func (l *DefaultLogger) _setup() error {
	logFile, err := os.OpenFile(filepath.Join(l.fullpath, "logger.log"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	l.logFile = logFile
	return nil
}

func (l *DefaultLogger) setup() {
	if err := l._setup(); err != nil {
		printAndExit(err.Error(), 1)
	}
}

func (l *DefaultLogger) _Log(message string, level LogLevel) error {
	if l.level > level {
		return nil
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

func (l *DefaultLogger) Log(message string, level LogLevel) {
	if err := l._Log(message, level); err != nil {
		printAndExit(err.Error(), 1)
	}
}

func (l *DefaultLogger) DEBUG(message string) {
	l.Log(message, DEBUG)
}

func (l *DefaultLogger) INFO(message string) {
	l.Log(message, INFO)
}

func (l *DefaultLogger) WARN(message string) {
	l.Log(message, WARN)
}

func (l *DefaultLogger) ERROR(message string) {
	l.Log(message, ERROR)
}
