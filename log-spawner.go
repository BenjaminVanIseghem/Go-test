package main

//All imports necessary for the logging chain
import (
	"math/rand"

	lkh "github.com/arriqaaq/logrus-kafka"
	log "github.com/sirupsen/logrus"
)

var infoLogger = createInfoLogger()

//Initiliazing of the logrus object
func init() {

}

func main() {
	for i := 0; i < 25; i++ {
		num := rand.Intn(100)
		infoLogger.Infof("Info number: %d", num)
	}
}

func createInfoLogger() *log.Entry {
	var tempLogger = log.New()
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

//EXAMPLES OF ALL SEVERITY LEVELS OF LOGRUS

//INFO
// log.WithFields(log.Fields{
// 	"service": "This is the service name",
// }).Info("This is an Info message ")

//DEBUG --> NOT USED
// log.WithFields(log.Fields{
// 	"service": "This is the service name",
// }).Debug("This is a Debug message ")

//TRACE --> NOT USED
// log.WithFields(log.Fields{
// 	"service": "This is the service name",
// }).Trace("This is a Trace message ")

//WARNING
// log.WithFields(log.Fields{
// 	"service": "This is the service name",
// }).Warn("This is a Warning message ")

//ERROR
// log.WithFields(log.Fields{
// 	"service": "This is the service name",
// }).Error("This is an Error message ")

//FATAL
// log.WithFields(log.Fields{
// 	"service": "This is the service name",
// }).Fatal("This is a Fatal message ")

//PANIC
// log.WithFields(log.Fields{
// 	"service": "This is the service name",
// }).Panic("This is a panic message ")

//END OF EXAMPLES
