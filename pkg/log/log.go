package log

import (
	"fmt"
	"net"
	"os"
	"runtime"

	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/jameshwc/Million-Singer/conf"
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
		Level:        logrus.TraceLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
		Formatter: &logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		},
	}
	conn, err := net.Dial("tcp", conf.LogConfig.LogStashAddr)
	if err != nil {
		logrus.Fatal(err)
	}
	hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{"type": "myappName"}))
	logger.Hooks.Add(hook)
}

// Debug output logs at debug level
func Debug(v ...interface{}) {
	logger.Debug(v...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Info output logs at info level
func Info(v ...interface{}) {
	logger.Info(v...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func InfoWithSource(v ...interface{}) {
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

func TraceIP(ipaddr, endpoint string) {
	logger.WithField("ip", ipaddr).WithField("endpoint", endpoint).Trace()
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
