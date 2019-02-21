package logger

import (
	lkh "github.com/arriqaaq/logrus-kafka"
	"github.com/sirupsen/logrus"
)

var kafkaLogger = createKafkaLogger()

//Info logs info messages
func Info(msg string) {
	kafkaLogger.WithFields(logrus.Fields{
		"id":           1,
		"info field 2": "some value",
		"info field 3": "another value",
	}).Info(msg)
}

//Warning logs warning messages
func Warning(msg string) {
	kafkaLogger.WithFields(logrus.Fields{
		"id":              2,
		"warning field 2": "some warning",
		"warning field 3": "another warning value",
	}).Warn(msg)
}

//Error logs error messages
func Error(msg string) {
	kafkaLogger.WithFields(logrus.Fields{
		"id":            3,
		"error field 2": "some error",
		"error field 3": "another error value",
	}).Error(msg)
}

func createKafkaLogger() *logrus.Logger {
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
		[]string{"192.168.99.222:9092"},
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