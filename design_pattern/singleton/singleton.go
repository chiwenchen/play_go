package main

import (
	"fmt"
	"sync"
)

type MyLogger struct {
	level string
}

func (m *MyLogger) Log(msg string) {
	fmt.Println(m.level, ": ", msg)
}

var logger *MyLogger

var once sync.Once

func getLoggerInstance() *MyLogger {
	once.Do(func() {
		fmt.Println("Creating new logger instance")
		logger = &MyLogger{"warning"}
	})

	// 下面是舊的作法，沒辦法確保在 gorouting之下，依然只建立一個
	// if logger == nil {
	// 	fmt.Println("Creating new logger instance")
	// 	logger = &MyLogger{"warning"}
	// }

	fmt.Println("Retriving logger instance")
	return logger
}

func (m *MyLogger) setLoggerLevel(level string) {
	m.level = level
}
