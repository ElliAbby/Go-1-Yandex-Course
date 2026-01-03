package main

import "fmt"

type Logger interface {
	Log(message string)
}

type LogLevel string

const (
	Error LogLevel = "Error"
	Info LogLevel = "Info"
)

type Log struct {
	Level LogLevel
}

func (l Log) Log(message string) {
	switch l.Level {
		case Error:
			fmt.Printf("ERROR: %s\n", message)
		case Info:
			fmt.Printf("INFO: %s\n", message)
	}
}

func main() {
	var errorLogger Logger
	var infoLogger Logger

	errorLogger = Log{Error}
	infoLogger = Log{Info}
	errorLogger.Log("Что-то пошло не так!")
	infoLogger.Log("Сервис успешно запущен.")
}
