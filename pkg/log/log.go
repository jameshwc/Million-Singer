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

var Logger *logrus.Logger

func Setup() {
	Logger = &logrus.Logger{
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
	if conf.LogConfig.IsEnabled {
		conn, err := net.Dial("tcp", conf.LogConfig.LogStashAddr)
		if err != nil {
			logrus.Fatal(err)
		}
		hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{"type": "cn_project"}))
		Logger.Hooks.Add(hook)
	}
}

// Debug output logs at debug level
func Debug(v ...interface{}) {
	Logger.Debug(v...)
}

func Debugf(format string, args ...interface{}) {
	Logger.Debugf(format, args...)
}

// Info output logs at info level
func Info(v ...interface{}) {
	Logger.Info(v...)
}

func Infof(format string, args ...interface{}) {
	Logger.Infof(format, args...)
}

func InfoWithSource(v ...interface{}) {
	Logger.WithField("source", getSource()).Info(v...)
}

// Warn output logs at warn level
func Warn(v ...interface{}) {
	Logger.Warn(v...)
}

func WarnWithSource(v ...interface{}) {
	Logger.WithField("source", getSource()).Error(v...)
}

// Error output logs at error level
func Error(v ...interface{}) {
	Logger.WithField("source", getSource()).Error(v...)
}

// Fatal output logs at fatal level
func Fatal(v ...interface{}) {
	Logger.WithField("source", getSource()).Fatal(v...)
}

func Fatalf(format string, args ...interface{}) {
	Logger.WithField("source", getSource()).Fatalf(format, args...)
}

func TraceIP(ipaddr, endpoint string) {
	Logger.WithField("ip", ipaddr).WithField("endpoint", endpoint).Trace()
}

// get source of the log output
func getSource() string {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if !ok {
		return "not available"
	}
	return fmt.Sprintf("%s:%d", file, line)
}
