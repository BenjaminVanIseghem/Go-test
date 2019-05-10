package main

import (
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	// w, err := os.Create("/Users/benjaminvaniseghem/Documents/promtail/logs/logfile2.log")
	// if err != nil {
	// 	panic(err)
	// }
	// logger := logrus.New()
	// logger.SetOutput(w)
	// for {
	// 	logger.Info("Info 2 message")
	// 	logger.Warn("Warning 2 message")
	// 	logger.Error("Error 2 message", errors.New("Error"))
	// 	time.Sleep(1200 * time.Millisecond)
	// }
	for {
		logrus.Info("Info message")
		logrus.Warn("Warning message")
		logrus.Error("Error message", errors.New("Error"))
		time.Sleep(1200 * time.Millisecond)
	}
}
