package log

import (
	"fmt"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

var (
	DefaultCallerDepth = 2
)

var logger *logrus.Logger

func Setup() {
	logger = &logrus.Logger{
		Out:          os.Stderr,
		Hooks:        make(logrus.LevelHooks),
		Level:        logrus.InfoLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
		Formatter: &logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		},
	}
}

// Debug output logs at debug level
func Debug(v ...interface{}) {
	logger.Debug(v...)
}

// Info output logs at info level
func Info(v ...interface{}) {
	logger.Info(v...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}
func InfoWithSource(v ...interface{}) {
	getSource()
	logger.WithField("source", getSource()).Info(v...)
}

// Warn output logs at warn level
func Warn(v ...interface{}) {
	logger.Warn(v...)
}
func WarnWithSource(v ...interface{}) {
	logger.WithField("source", getSource()).Error(v...)
}

// Error output logs at error level
func Error(v ...interface{}) {
	logger.WithField("source", getSource()).Error(v...)
}

// Fatal output logs at fatal level
func Fatal(v ...interface{}) {
	logger.WithField("source", getSource()).Fatal(v...)
}

func Fatalf(format string, args ...interface{}) {
	logger.WithField("source", getSource()).Fatalf(format, args...)
}

// get source of the log output
func getSource() (source string) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		source = fmt.Sprintf("%s:%d", file, line)
	} else {
		source = "not available"
	}
	return
}
