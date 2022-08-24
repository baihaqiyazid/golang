package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T)  {
	log := logrus.New()

	log.Println("Hello World")
	fmt.Println("Hello World")
}

func TestLevelLogger(t *testing.T)  {
	log := logrus.New()
	log.SetLevel(logrus.TraceLevel)

	log.Trace("This is Trace")
	log.Debug("This is Debug")
	log.Info("This is Info")
	log.Warn("This is Warning")
	log.Error("This is Error")
	log.Fatal("This is Fatal")
	log.Panic("This is Panic")
}

func TestOutLogger(t *testing.T)  {
	log := logrus.New()
	log.SetLevel(logrus.TraceLevel)

	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666 )
	log.SetOutput(file)

	log.Trace("This is Trace")
	log.Debug("This is Debug")
	log.Info("This is Info")
	log.Warn("This is Warning")
	log.Error("This is Error")
}


func TestFormatterLogger(t *testing.T)  {
	log := logrus.New()
	log.SetLevel(logrus.TraceLevel)

	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666 )
	log.SetOutput(file)

	log.SetFormatter(&logrus.JSONFormatter{})

	log.Trace("This is Trace")
	log.Debug("This is Debug")
	log.Info("This is Info")
	log.Warn("This is Warning")
	log.Error("This is Error")
}

func TestFieldsLogger(t *testing.T)  {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	log.WithFields(logrus.Fields{
		"username": "Baihaqi123",
		"name": "Muhammad Yazid Baihaqi",
	}).Info("test field")
}
