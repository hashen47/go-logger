package logger

import (
	"fmt"
	"os"
	"time"
)

func printAndExit(data string, exitCode int) {
	fmt.Println(data)
	os.Exit(exitCode)
}

func getFormatterTokenType(toketypeStr string) (FormatterTokenType, bool) {
	switch toketypeStr {
	case "datetime":
		return DATETIME, true
	case "level":
		return LEVEL, true
	case "message":
		return MESSAGE, true
	}
	return FormatterTokenType(-1), false
}

func tokenizeFormatString(formatString string) (*[]FormatterToken, error) {
	tokens := make([]FormatterToken, 0)

	isOnlyTEXTTokenContains := true
	i1, i2, i3 := 0, -1, -1

	for i := 0; i < len(formatString); i++ {
		ch := formatString[i]

		if ch == '%' {
			if i2 == -1 {
				i2 = i
				continue
			}

			i3 = i
			tokenType, isFind := getFormatterTokenType(formatString[i2+1 : i3])
			if !isFind {
				return nil, &InvalidFormatStringError{formatString, "formatter token type can only be %message%, %datetime% or %level%"}
			}

			if i1 < i2 {
				tokens = append(tokens, FormatterToken{
					ttype: TEXT,
					data:  formatString[i1:i2],
				})
			}

			tokens = append(tokens, FormatterToken{
				ttype: tokenType,
				data:  "",
			})

			isOnlyTEXTTokenContains = false

			i1 = i3 + 1
			i2, i3 = -1, -1
		}
	}

	if i2 != -1 {
		return nil, &InvalidFormatStringError{formatString, "you only can use % sign with formatter token types"}
	}

	if i1 < len(formatString) {
		tokens = append(tokens, FormatterToken{
			ttype: TEXT,
			data:  formatString[i1:len(formatString)],
		})
	}

	if len(tokens) == 0 {
		return nil, &InvalidFormatStringError{formatString, "format string cannot be empty"}
	}

	if isOnlyTEXTTokenContains {
		return nil, &InvalidFormatStringError{formatString, "at least should contains one other formatter token type (%message%, %level% or %datetime%)"}
	}

	return &tokens, nil
}

func validateLogLevel(level LogLevel) error {
	switch level {
	case DEBUG, INFO, WARN, ERROR:
		return nil
	default:
		return &InvalidLogLevelError{level}
	}
	return nil
}

func validateLoggerType(loggerType LoggerType) error {
	switch loggerType {
	case DEFAULT_TYPE, LOG_LEVEL_ORIENTED_TYPE, DATE_ORIENTED_TYPE:
		return nil
	default:
		return &InvalidLoggerTypeError{loggerType}
	}
	return nil
}

func isPathExists(path string) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}
	return nil
}

func getLogLevelStr(level LogLevel) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	}
	return ""
}

func getFormatterTokenValue(token FormatterToken, userInput string, level LogLevel) (string, error) {
	switch token.ttype {
	case TEXT:
		return token.data, nil
	case MESSAGE:
		return userInput, nil
	case DATETIME:
		return time.Now().Format(time.DateTime), nil
	case LEVEL:
		return getLogLevelStr(level), nil
	}
	return "", &InvalidFormatterTokenTypeError{token.ttype}
}
