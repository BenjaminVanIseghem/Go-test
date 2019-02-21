package logger

import (
	lkh "github.com/arriqaaq/logrus-kafka"
	log "github.com/sirupsen/logrus"
)

//CreateInfoLogger is a logger with fields specific to the level INFO
func CreateInfoLogger() *log.Entry {
	var tempLogger = log.New()
	//SetReportCaller prints the calling method where the log happens
	tempLogger.SetReportCaller(true)
	//Create hook using logrus-kafka library
	//Params are id, severity levels that will be passed, formatter, kafka broker, kafka topic
	infoHook, err := lkh.NewKafkaSyncHook(
		"kh",
		[]log.Level{log.InfoLevel, log.WarnLevel, log.ErrorLevel},
		&log.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		},
		[]string{"192.168.99.222:9092"},
		"info")
	//Error handling for the creating of the hook
	if err != nil {
		log.Error("Problem with kafka hook")
	}

	//Add hook to temporary logger object
	tempLogger.Hooks.Add(infoHook)

	//Create infologger
	infoLogger := tempLogger.WithFields(log.Fields{
		"chain_id": "Which part of the chain",
		"service":  "service name",
		"event_id": 1,
	})
	//Debug log
	log.Debug("created info logger")

	return infoLogger
}
