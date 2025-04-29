package logger

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type LoggerType int

const (
	DEFAULT_TYPE            LoggerType = iota // all logs stored in one file
	LOG_LEVEL_ORIENTED_TYPE                   // each log level have different log file
	DATE_ORIENTED_TYPE                        // each date different log file will created
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

type LoggerI interface {
	setup()
	GetFullPath() string
	Log(message string, level LogLevel)
	DEBUG(message string)
	INFO(message string)
	WARN(message string)
	ERROR(message string)
}

type Logger struct {
	fullpath  string
	level     LogLevel
	formatter *Formatter
}

func _NewLogger(basepath string, dir string, level LogLevel, ltype LoggerType, formatter *Formatter) (LoggerI, error) {
	var logger LoggerI

	basepath = strings.Trim(basepath, " ")
	dir = strings.Trim(dir, " ")

	if err := isPathExists(basepath); err != nil {
		return nil, err
	}

	if dir == "" {
		return nil, &LoggerDirNameEmptyError{}
	}

	fullpath := filepath.Join(basepath, dir)
	if err := isPathExists(fullpath); err != nil {
		if err := os.MkdirAll(fullpath, 0755); err != nil {
			return nil, err
		}
	}

	if err := validateLogLevel(level); err != nil {
		os.RemoveAll(fullpath)
		return nil, err
	}

	if formatter == nil {
		os.RemoveAll(fullpath)
		return nil, &InvalidFormatterError{"formatter cannot be nil"}
	}

	switch ltype {
	case DEFAULT_TYPE:
		logger = &DefaultLogger{
			Logger{
				fullpath,
				level,
				formatter,
			},
			nil,
			&sync.Mutex{},
		}
	case LOG_LEVEL_ORIENTED_TYPE:
		logger = &LevelOrientedLogger{
			Logger{
				fullpath,
				level,
				formatter,
			},
			nil,
			nil,
			nil,
			nil,
			&sync.Mutex{},
		}
	case DATE_ORIENTED_TYPE:
		logger = &DateOrientedLogger{
			Logger{
				fullpath,
				level,
				formatter,
			},
			nil,
			"",
			&sync.Mutex{},
		}
	default:
		os.RemoveAll(fullpath)
		return nil, &InvalidLoggerTypeError{ltype}
	}

	logger.setup()

	return logger, nil
}

func NewLogger(basepath string, dir string, level LogLevel, ltype LoggerType, formatter *Formatter) LoggerI {
	logger, err := _NewLogger(basepath, dir, level, ltype, formatter)
	if err != nil {
		printAndExit(err.Error(), 1)
	}
	return logger
}

func (l *Logger) SetLevel(level LogLevel) {
	if err := validateLogLevel(level); err != nil {
		printAndExit(err.Error(), 1)
	}
	l.level = level
}

func (l *Logger) GetFullPath() string {
	return l.fullpath
}
