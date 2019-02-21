package main

//All imports necessary for the logging chain
import (
	"log-spawner/logger"
)

var m1 = make(map[string]interface{})

//Initiliazing of the logrus object
func init() {
	m1["fixed key 1"] = "string 1"
	m1["fixed key 2"] = "string 2"
	m1["fixed key 3"] = 8
}

func main() {
	// for i := 0; i < 25; i++ {
	// 	logger.Info("Info message")
	// 	logger.Warning("Warning message")
	// 	logger.Error("Error message")
	// }

	m2 := make(map[string]interface{})

	m2["key1"] = "value1"
	m2["key2"] = "value2"
	m2["key3"] = "value3"
	m2["key4"] = 6

	logger.Info("info message", m1)
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
