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
	AppLog   = "app"
	ErrorLog = "error"
	HttpLog  = "http"
	DbLog    = "database"
)

func init() {
	Logger = GetLogger()
	Logger.Set(HttpLog, log.New(Logger.CreateLogFileWriter("access.log"), "[http] ", log.LstdFlags))
	Logger.Set(AppLog, log.New(Logger.CreateLogFileWriter("app.log", os.Stdout), "[app] ", log.LstdFlags))
	Logger.Set(DbLog, log.New(Logger.CreateLogFileWriter("database.log"), "[database] ", log.LstdFlags))
	Logger.Set(ErrorLog, log.New(Logger.CreateLogFileWriter("error.log", os.Stderr), "[database] ", log.LstdFlags))
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

func (l *logger) CreateLogFileWriter(filename string, writers ...io.Writer) io.Writer {
	err := os.MkdirAll("./var/log/", 0666)
	if err != nil {
		log.Panic(err)
	}

	file, err := os.OpenFile("./var/log/"+filename, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Panic(err)
	}

	writer := io.MultiWriter(append(writers, file)...)

	return writer
}
