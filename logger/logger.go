package logger

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func Info(msg string) {
	log.Info(msg)
}

func Infof(format string, v ...any) {
	log.Infof(format, v...)
}

func Error(msg string) {
	log.Error(msg)
}

func Errorf(format string, v ...any) {
	log.Errorf(format, v...)
}

func Debug(msg string) {
	log.Debug(msg)
}

func Debugf(format string, v ...any) {
	log.Debugf(format, v...)
}

func Warn(msg string) {
	log.Warn(msg)
}

func Warnf(format string, v ...any) {
	log.Warnf(format, v...)
}
