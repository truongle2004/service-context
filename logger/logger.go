package logger

import (
	"fmt"
	"log"
)

func Info(msg string) {
	log.Println("[INFO]", msg)
}

func Infof(format string, v ...any) {
	log.Println("[INFO]", fmt.Sprintf(format, v...))
}

func Error(msg string) {
	log.Println("[ERROR]", msg)
}

func Errorf(format string, v ...any) {
	log.Println("[ERROR]", fmt.Sprintf(format, v...))
}

func Debug(msg string) {
	log.Println("[DEBUG]", msg)
}

func Debugf(format string, v ...any) {
	log.Println("[DEBUG]", fmt.Sprintf(format, v...))
}

func Warn(msg string) {
	log.Println("[WARN]", msg)
}

func Warnf(format string, v ...any) {
	log.Println("[WARN]", fmt.Sprintf(format, v...))
}
