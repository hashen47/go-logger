package logger

import (
	"os"
	"path/filepath"
)

type LevelOrientedLogger struct {
	Logger
	debugLogFile *os.File
	infoLogFile  *os.File
	warnLogFile  *os.File
	errorLogFile *os.File
}

func (l *LevelOrientedLogger) _setup() error {
	debugLogFile, err := os.OpenFile(filepath.Join(l.fullpath, "debug.log"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	l.debugLogFile = debugLogFile

	infoLogFile, err := os.OpenFile(filepath.Join(l.fullpath, "info.log"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	l.infoLogFile = infoLogFile

	warnLogFile, err := os.OpenFile(filepath.Join(l.fullpath, "warn.log"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	l.warnLogFile = warnLogFile

	errorLogFile, err := os.OpenFile(filepath.Join(l.fullpath, "error.log"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	l.errorLogFile = errorLogFile

	return nil
}

func (l *LevelOrientedLogger) setup() {
	if err := l._setup(); err != nil {
		printAndExit(err.Error(), 1)
	}
}

func (l *LevelOrientedLogger) _Log(message string, level LogLevel) error {
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

	output += "\n"

	switch level {
	case DEBUG:
		if _, err := l.debugLogFile.Write([]byte(output)); err != nil {
			return err
		}
	case INFO:
		if _, err := l.infoLogFile.Write([]byte(output)); err != nil {
			return err
		}
	case WARN:
		if _, err := l.warnLogFile.Write([]byte(output)); err != nil {
			return err
		}
	case ERROR:
		if _, err := l.errorLogFile.Write([]byte(output)); err != nil {
			return err
		}
	default:
		return &InvalidLogLevelError{level}
	}

	return nil
}

func (l *LevelOrientedLogger) Log(message string, level LogLevel) {
	if err := l._Log(message, level); err != nil {
		printAndExit(err.Error(), 1)
	}
}

func (l *LevelOrientedLogger) DEBUG(message string) {
	l.Log(message, DEBUG)
}

func (l *LevelOrientedLogger) INFO(message string) {
	l.Log(message, INFO)
}

func (l *LevelOrientedLogger) WARN(message string) {
	l.Log(message, WARN)
}

func (l *LevelOrientedLogger) ERROR(message string) {
	l.Log(message, ERROR)
}
