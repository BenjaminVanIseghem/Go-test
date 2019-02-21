package main

//All imports necessary for the logging chain
import (
	"log-spawner/logger"
	"math/rand"
)

var infoLogger = logger.CreateInfoLogger()

//Initiliazing of the logrus object
func init() {

}

func main() {
	for i := 0; i < 25; i++ {
		num := rand.Intn(100)
		infoLogger.Infof("Info number: %d", num)
	}
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
