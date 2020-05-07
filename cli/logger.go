package cli

import (
	"fmt"
	"io"
	"log"
	"os"
)

type logger struct {
	loggers map[string]interface{}
}

var Logger *logger

const (
	AppLog  = "app"
	HttpLog = "http"
)

func init() {
	Logger = GetLogger()
	Logger.Set(HttpLog, log.New(Logger.CreateLogFile("access.log"), "[http] ", log.LstdFlags))
	Logger.Set(AppLog, log.New(os.Stdout, "[app] ", log.LstdFlags))
}

func GetLogger() (l *logger) {
	l = &logger{loggers: make(map[string]interface{})}

	return
}

func (l *logger) Get(key string) (instance interface{}) {
	var exists bool
	instance, exists = l.loggers[key]

	if !exists {
		panic(fmt.Sprintf(`Logger for "%s" does not exists.`, key))
	}

	return
}

func (l *logger) Set(key string, instance interface{}) {
	l.loggers[key] = instance
}

func (l *logger) CreateLogFile(filename string) io.Writer {
	err := os.MkdirAll("./var/log/", 0666)
	if err != nil {
		log.Panic(err)
	}

	file, err := os.OpenFile("./var/log/"+filename, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Panic(err)
	}

	writer := io.MultiWriter(file)

	return writer
}