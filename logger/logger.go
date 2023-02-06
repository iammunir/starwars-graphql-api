package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.elastic.co/ecslogrus"
)

type Logger interface {
	Trace(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
}

func InitLogger() Logger {

	log := logrus.New()

	log.SetFormatter(&ecslogrus.Formatter{})
	log.ReportCaller = true

	logLevel := viper.GetUint16("LOG_LEVEL")

	log.SetLevel(logrus.Level(logLevel))

	return log
}
