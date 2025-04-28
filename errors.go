package logger

import (
	"fmt"
)

type InvalidLogLevelError struct {
	level LogLevel
}

func (e *InvalidLogLevelError) Error() string {
	return fmt.Sprintf("\n------------------\nINVALID LOG LEVEL\nLevel: %d\n------------------\n", int(e.level))
}

type InvalidLoggerTypeError struct {
	loggerType LoggerType
}

func (e *InvalidLoggerTypeError) Error() string {
	return fmt.Sprintf("\n------------------\nINVALID LOGGER TYPE\nType: %d\n------------------\n", int(e.loggerType))
}

type InvalidFormatStringError struct {
	formatStr string
	extraMsg  string
}

func (e *InvalidFormatStringError) Error() string {
	return fmt.Sprintf("\n------------------\nINVALID FORMAT STRING\nFormat String: %s\nMessage: %s\n------------------\n", e.formatStr, e.extraMsg)
}

type InvalidFormatterError struct {
	msg string
}

func (e *InvalidFormatterError) Error() string {
	return fmt.Sprintf("\n------------------\nINVALID FORMATTER ERROR\nMessage: %s\n------------------\n", e.msg)
}

type LoggerDirNameEmptyError struct{}

func (e *LoggerDirNameEmptyError) Error() string {
	return fmt.Sprintf("\n------------------\nLOGGER DIR NAME CANNOT BE EMPTY\n------------------\n")
}

type InvalidFormatterTokenTypeError struct {
	ttype FormatterTokenType
}

func (e *InvalidFormatterTokenTypeError) Error() string {
	return fmt.Sprintf("\n------------------\nINVALID FORMATTER TOKEN TYPE\nToken Type: %d\n------------------\n", int(e.ttype))
}
