package main

import (
	"errors"

	"github.com/sirupsen/logrus"
)

func main() {
	for {
		logrus.Info("Info message")
		logrus.Warn("Warning message")
		logrus.Error("Error message", errors.New("Error"))
	}
}
