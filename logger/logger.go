package logger

import (
	lkh "github.com/arriqaaq/logrus-kafka"
	"github.com/sirupsen/logrus"
)

var infoLogger = createInfoLogger()
var warnLogger = createWarningLogger()
var errLogger = createErrorLogger()

//Info logs info messages
func Info(msg string, fields map[string]interface{}) {
	//Convert fields parameter to logrus fields type
	fields = logrus.Fields(fields)

	//Log the given fields and message via logrus
	infoLogger.WithFields(
		fields,
	).Info(msg)
}

//Warning logs warning messages
func Warning(msg string, fields map[string]interface{}) {
	//Convert fields parameter to logrus fields type
	fields = logrus.Fields(fields)

	//Log the given fields and message via logrus
	warnLogger.WithFields(
		fields,
	).Warning(msg)
}

//Error logs error messages
func Error(msg string, fields map[string]interface{}) {
	//Convert fields parameter to logrus fields type
	fields = logrus.Fields(fields)

	//Log the given fields and message via logrus
	errLogger.WithFields(
		fields,
	).Error(msg)
}

//Create a logger that uses a hook to post logs to kafka on the "info" topic
func createInfoLogger() *logrus.Logger {
	var tempLogger = logrus.New()
	tempLogger.SetReportCaller(true)

	//Create hook using logrus-kafka library
	//Params are id, severity levels that will be passed, formatter, kafka broker, kafka topic
	hook, err := lkh.NewKafkaSyncHook(
		"kh",
		[]logrus.Level{logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel},
		&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		},
		[]string{"localhost:9092"},
		"info")
	//Error handling for the creating of the hook
	if err != nil {
		logrus.Error("Problem with kafka hook")
	}
	//Add hook to the tempLogger
	tempLogger.Hooks.Add(hook)

	//Debug log
	logrus.Debug("created logger")

	return tempLogger
}

//Create a logger that uses a hook to post logs to kafka on the "warning" topic
func createWarningLogger() *logrus.Logger {
	var tempLogger = logrus.New()
	tempLogger.SetReportCaller(true)

	//Create hook using logrus-kafka library
	//Params are id, severity levels that will be passed, formatter, kafka broker, kafka topic
	hook, err := lkh.NewKafkaSyncHook(
		"kh",
		[]logrus.Level{logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel},
		&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		},
		[]string{"localhost:9092"},
		"warning")
	//Error handling for the creating of the hook
	if err != nil {
		logrus.Error("Problem with kafka hook")
	}
	//Add hook to the tempLogger
	tempLogger.Hooks.Add(hook)

	//Debug log
	logrus.Debug("created logger")

	return tempLogger
}

//Create a logger that uses a hook to post logs to kafka on the "error" topic
func createErrorLogger() *logrus.Logger {
	var tempLogger = logrus.New()
	tempLogger.SetReportCaller(true)

	//Create hook using logrus-kafka library
	//Params are id, severity levels that will be passed, formatter, kafka broker, kafka topic
	hook, err := lkh.NewKafkaSyncHook(
		"kh",
		[]logrus.Level{logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel},
		&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		},
		[]string{"localhost:9092"},
		"error")
	//Error handling for the creating of the hook
	if err != nil {
		logrus.Error("Problem with kafka hook")
	}
	//Add hook to the tempLogger
	tempLogger.Hooks.Add(hook)

	//Debug log
	logrus.Debug("created logger")

	return tempLogger
}
