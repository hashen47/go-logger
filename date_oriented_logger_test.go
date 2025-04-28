package logger

import (
	"bufio"
	"errors"
	"os"
	"testing"
	"time"
)

func TestDateOrientedNewLogger(t *testing.T) {
	type LogMessage struct {
		msg    string
		level  LogLevel
		expect string
	}

	type Testcase struct {
		basepath    string
		dir         string
		level       LogLevel
		ltype       LoggerType
		formatter   *Formatter
		err         error
		logMessages []LogMessage
		curTime     string
	}

	curDir, _ := os.Getwd()

	testcases := []Testcase{
		{
			basepath:  curDir,
			dir:       "logs-date-oriented-logger",
			level:     DEBUG,
			ltype:     DATE_ORIENTED_TYPE,
			formatter: NewFormatter("[%level%]: %message%"),
			err:       nil,
			logMessages: []LogMessage{
				{
					"this is first debug msg",
					DEBUG,
					"[DEBUG]: this is first debug msg",
				},
				{
					"this is first info msg",
					INFO,
					"[INFO]: this is first info msg",
				},
				{
					"this is first warn msg",
					WARN,
					"[WARN]: this is first warn msg",
				},
				{
					"this is first error msg",
					ERROR,
					"[ERROR]: this is first error msg",
				},
			},
			curTime: "2028-02-22",
		},
		{
			basepath:  curDir,
			dir:       "somedir",
			level:     DEBUG,
			ltype:     DATE_ORIENTED_TYPE,
			formatter: NewFormatter("%message%"),
			err:       nil,
			logMessages: []LogMessage{
				{
					"this is first debug msg",
					DEBUG,
					"this is first debug msg",
				},
				{
					"this is first info msg",
					INFO,
					"this is first info msg",
				},
				{
					"this is first warn msg",
					WARN,
					"this is first warn msg",
				},
				{
					"this is first error msg",
					ERROR,
					"this is first error msg",
				},
			},
			curTime: "2028-02-22",
		},
		{
			basepath:  curDir,
			dir:       "logdir",
			level:     DEBUG,
			ltype:     DATE_ORIENTED_TYPE,
			formatter: NewFormatter("[%level%]"),
			err:       nil,
			logMessages: []LogMessage{
				{
					"this is first debug msg",
					DEBUG,
					"[DEBUG]",
				},
				{
					"this is first info msg",
					INFO,
					"[INFO]",
				},
				{
					"this is first warn msg",
					WARN,
					"[WARN]",
				},
				{
					"this is first error msg",
					ERROR,
					"[ERROR]",
				},
			},
			curTime: "2024-01-28",
		},
		{
			basepath:  curDir,
			dir:       "loggins",
			level:     DEBUG,
			ltype:     DATE_ORIENTED_TYPE,
			formatter: NewFormatter("[%level%]: %message% %message%"),
			err:       nil,
			logMessages: []LogMessage{
				{
					"this is first debug msg",
					DEBUG,
					"[DEBUG]: this is first debug msg this is first debug msg",
				},
				{
					"this is second debug msg",
					DEBUG,
					"[DEBUG]: this is second debug msg this is second debug msg",
				},
				{
					"",
					DEBUG,
					"[DEBUG]:  ",
				},
				{
					"this is first info msg",
					INFO,
					"[INFO]: this is first info msg this is first info msg",
				},
				{
					"this is first warn msg",
					WARN,
					"[WARN]: this is first warn msg this is first warn msg",
				},
				{
					"this is first error msg",
					ERROR,
					"[ERROR]: this is first error msg this is first error msg",
				},
			},
			curTime: "2023-10-15",
		},
		{
			basepath:  curDir,
			dir:       "log-dir",
			level:     DEBUG,
			ltype:     DATE_ORIENTED_TYPE,
			formatter: NewFormatter("%message% %message%"),
			err:       nil,
			logMessages: []LogMessage{
				{
					"this is first debug msg",
					DEBUG,
					"this is first debug msg this is first debug msg",
				},
				{
					"this is second debug msg",
					DEBUG,
					"this is second debug msg this is second debug msg",
				},
				{
					"",
					DEBUG,
					" ",
				},
				{
					"this is first info msg",
					INFO,
					"this is first info msg this is first info msg",
				},
				{
					"this is first warn msg",
					WARN,
					"this is first warn msg this is first warn msg",
				},
				{
					"this is first error msg",
					ERROR,
					"this is first error msg this is first error msg",
				},
			},
			curTime: "2026-11-08",
		},
		{
			basepath:    "/something-that-not-exists",
			dir:         "log-dir",
			level:       DEBUG,
			ltype:       DATE_ORIENTED_TYPE,
			formatter:   NewFormatter("%message% %message%"),
			err:         errors.New("stat " + "/something-that-not-exists: no such file or directory"),
			logMessages: []LogMessage{},
			curTime:     "2026-11-08",
		},
		{
			basepath:    "/not-exists-directory",
			dir:         "log-dir",
			level:       DEBUG,
			ltype:       DATE_ORIENTED_TYPE,
			formatter:   NewFormatter("%message% %message%"),
			err:         errors.New("stat " + "/not-exists-directory: no such file or directory"),
			logMessages: []LogMessage{},
			curTime:     "2026-11-08",
		},
		{
			basepath:    curDir,
			dir:         "",
			level:       DEBUG,
			ltype:       DATE_ORIENTED_TYPE,
			formatter:   NewFormatter("%message% %message%"),
			err:         &LoggerDirNameEmptyError{},
			logMessages: []LogMessage{},
			curTime:     "2026-11-08",
		},
		{
			basepath:    curDir,
			dir:         "  ",
			level:       DEBUG,
			ltype:       DATE_ORIENTED_TYPE,
			formatter:   NewFormatter("%message% %message%"),
			err:         &LoggerDirNameEmptyError{},
			logMessages: []LogMessage{},
			curTime:     "2026-11-08",
		},
		{
			basepath:    curDir + "/something-that-not-exists",
			dir:         "log-dir",
			level:       DEBUG,
			ltype:       DATE_ORIENTED_TYPE,
			formatter:   NewFormatter("%message% %message%"),
			err:         errors.New("stat " + curDir + "/something-that-not-exists: no such file or directory"),
			logMessages: []LogMessage{},
			curTime:     "2026-11-08",
		},
		{
			basepath:    curDir + "/not-a-directory",
			dir:         "log-dir",
			level:       DEBUG,
			ltype:       DATE_ORIENTED_TYPE,
			formatter:   NewFormatter("%message% %message%"),
			err:         errors.New("stat " + curDir + "/not-a-directory: no such file or directory"),
			logMessages: []LogMessage{},
			curTime:     "2026-11-08",
		},
		{
			basepath:    curDir,
			dir:         "log-dir",
			level:       LogLevel(23),
			ltype:       DATE_ORIENTED_TYPE,
			formatter:   NewFormatter("%message% %message%"),
			err:         &InvalidLogLevelError{23},
			logMessages: []LogMessage{},
			curTime:     "2026-11-08",
		},
		{
			basepath:    curDir,
			dir:         "log-dir-test",
			level:       LogLevel(-21),
			ltype:       DATE_ORIENTED_TYPE,
			formatter:   NewFormatter("%message% %message%"),
			err:         &InvalidLogLevelError{-21},
			logMessages: []LogMessage{},
			curTime:     "2026-11-08",
		},
		{
			basepath:    curDir,
			dir:         "log-dir",
			level:       LogLevel(33),
			ltype:       DATE_ORIENTED_TYPE,
			formatter:   NewFormatter("%message% %message%"),
			err:         &InvalidLogLevelError{33},
			logMessages: []LogMessage{},
			curTime:     "2026-11-08",
		},
		{
			basepath:    curDir,
			dir:         "log-dir",
			level:       DEBUG,
			ltype:       DATE_ORIENTED_TYPE,
			formatter:   nil,
			err:         &InvalidFormatterError{"formatter cannot be nil"},
			logMessages: []LogMessage{},
			curTime:     "2026-11-08",
		},
		{
			basepath:  curDir,
			dir:       "loggins",
			level:     INFO,
			ltype:     DATE_ORIENTED_TYPE,
			formatter: NewFormatter("[%level%]: %message% %message%"),
			err:       nil,
			logMessages: []LogMessage{
				{
					"this is first debug msg",
					DEBUG,
					"",
				},
				{
					"this is second debug msg",
					DEBUG,
					"",
				},
				{
					"",
					DEBUG,
					"",
				},
				{
					"this is first info msg",
					INFO,
					"[INFO]: this is first info msg this is first info msg",
				},
				{
					"this is first warn msg",
					WARN,
					"[WARN]: this is first warn msg this is first warn msg",
				},
				{
					"this is first error msg",
					ERROR,
					"[ERROR]: this is first error msg this is first error msg",
				},
			},
			curTime: "2029-04-14",
		},
		{
			basepath:  curDir,
			dir:       "loggins",
			level:     WARN,
			ltype:     DATE_ORIENTED_TYPE,
			formatter: NewFormatter("[%level%]: %message% %message%"),
			err:       nil,
			logMessages: []LogMessage{
				{
					"this is first debug msg",
					DEBUG,
					"",
				},
				{
					"this is second debug msg",
					DEBUG,
					"",
				},
				{
					"",
					DEBUG,
					"",
				},
				{
					"this is first info msg",
					INFO,
					"",
				},
				{
					"this is first warn msg",
					WARN,
					"[WARN]: this is first warn msg this is first warn msg",
				},
				{
					"this is first error msg",
					ERROR,
					"[ERROR]: this is first error msg this is first error msg",
				},
			},
			curTime: "2019-07-11",
		},
		{
			basepath:  curDir,
			dir:       "loggins",
			level:     ERROR,
			ltype:     DATE_ORIENTED_TYPE,
			formatter: NewFormatter("[%level%]: %message% %message%"),
			err:       nil,
			logMessages: []LogMessage{
				{
					"this is first debug msg",
					DEBUG,
					"",
				},
				{
					"this is second debug msg",
					DEBUG,
					"",
				},
				{
					"",
					DEBUG,
					"",
				},
				{
					"this is first info msg",
					INFO,
					"",
				},
				{
					"this is first warn msg",
					WARN,
					"",
				},
				{
					"this is first error msg",
					ERROR,
					"[ERROR]: this is first error msg this is first error msg",
				},
			},
			curTime: "2018-07-21",
		},
		{
			basepath:    curDir,
			dir:         "log-dir",
			level:       DEBUG,
			ltype:       LoggerType(-1),
			formatter:   NewFormatter("%message% %message%"),
			err:         &InvalidLoggerTypeError{LoggerType(-1)},
			logMessages: []LogMessage{},
			curTime:     "2018-07-21",
		},
		{
			basepath:    curDir,
			dir:         "log-dir",
			level:       DEBUG,
			ltype:       LoggerType(24),
			formatter:   NewFormatter("%message% %message%"),
			err:         &InvalidLoggerTypeError{LoggerType(24)},
			logMessages: []LogMessage{},
			curTime:     "2018-07-21",
		},
		{
			basepath:    curDir,
			dir:         "log-dir",
			level:       DEBUG,
			ltype:       LoggerType(10),
			formatter:   NewFormatter("%message% %message%"),
			err:         &InvalidLoggerTypeError{LoggerType(10)},
			logMessages: []LogMessage{},
			curTime:     "2018-07-21",
		},
	}

	for _, tc := range testcases {
		logger, err := _NewLogger(tc.basepath, tc.dir, tc.level, tc.ltype, tc.formatter)
		if err != nil && tc.err == nil ||
			err == nil && tc.err != nil {
			t.Fatalf("\nUNEXPECTED LOGGER ERROR\nEXPECT: %v\nREAL: %v\n", tc.err, err)
		}

		if err != nil {
			if err.Error() != tc.err.Error() {
				t.Fatalf("\nUNEXPECTED LOGGER ERROR\nEXPECT: %v\nREAL: %v\n", tc.err, err)
			}
			continue
		}

		curTime, _ := time.Parse(time.DateOnly, tc.curTime)
		dateOrientedLogger := logger.(*DateOrientedLogger)
		dateOrientedLogger._setup(curTime)

		for _, logMessage := range tc.logMessages {
			switch logMessage.level {
			case DEBUG:
				err = dateOrientedLogger._Log(logMessage.msg, DEBUG, curTime)
			case INFO:
				err = dateOrientedLogger._Log(logMessage.msg, INFO, curTime)
			case WARN:
				err = dateOrientedLogger._Log(logMessage.msg, WARN, curTime)
			case ERROR:
				err = dateOrientedLogger._Log(logMessage.msg, ERROR, curTime)
			}
		}

		if err != nil {
			t.Fatalf("\nLOGGING ERROR\nEXPECT: %v\nREAL: %v\n", nil, err)
		}

		logFile, err := os.OpenFile(dateOrientedLogger.logFile.Name(), os.O_RDONLY, 0644)
		if err != nil {
			t.Fatalf("\nFILE OPEN ERROR\nEXPECT: %v\nREAL: %v\n", nil, err)
		}
		defer logFile.Close()

		if logFile.Name() != dateOrientedLogger.logFile.Name() {
			t.Fatalf("\nFILE NAME ERROR\nEXPECT: %v\nREAL: %v\n", dateOrientedLogger.logFile.Name(), logFile.Name())
		}

		logFileReader := bufio.NewReader(logFile)

		for _, logMessage := range tc.logMessages {
			if logMessage.expect == "" {
				continue
			}

			content, _ := logFileReader.ReadString('\n')
			content = content[:len(content)-1]
			if content != logMessage.expect {
				t.Fatalf("\nLOGGER CONTENT ERROR\nEXPECT: %v\nREAL: %v\n", logMessage.expect, content)
			}
		}

		if err := os.RemoveAll(dateOrientedLogger.fullpath); err != nil {
			if err != nil {
				t.Fatalf("\nFILE DELETION ERROR\nEXPECT: %v\nREAL: %v\n", nil, err)
			}
		}
	}
}
