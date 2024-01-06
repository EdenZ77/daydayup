package main

import (
	"fmt"
)

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
)

// Logger 接口仅需要定义 Log 方法
type Logger interface {
	Log(level Level, message string)
}

// BaseLogger 结构体包含一个抽象的 doLog 方法
type BaseLogger struct {
	name string
}

func NewBaseLogger(name string) *BaseLogger {
	return &BaseLogger{name: name}
}

// Log 调用 doLog 方法
func (l *BaseLogger) Log(level Level, message string) {
	// 此处抽象调用，具体实现由子类提供
	l.doLog(level, message)
}

// doLog 作为抽象方法，具体实现由子类提供
func (l *BaseLogger) doLog(level Level, message string) {
	// 抽象方法，无实现
}

// FileLogger 结构体嵌入 BaseLogger，并提供 doLog 的具体实现
type FileLogger struct {
	*BaseLogger
}

func NewFileLogger(name string) *FileLogger {
	return &FileLogger{
		BaseLogger: NewBaseLogger(name),
	}
}

// doLog 在 FileLogger 中的具体实现
func (f *FileLogger) doLog(level Level, message string) {
	fmt.Printf("[%s] [File] [%v]: %s\n", f.name, level, message)
}

// MessageQueueLogger 结构体嵌入 BaseLogger，并提供 doLog 的具体实现
type MessageQueueLogger struct {
	*BaseLogger
}

func NewMessageQueueLogger(name string) *MessageQueueLogger {
	return &MessageQueueLogger{
		BaseLogger: NewBaseLogger(name),
	}
}

// doLog 在 MessageQueueLogger 中的具体实现
func (m *MessageQueueLogger) doLog(level Level, message string) {
	fmt.Printf("[%s] [MessageQueue] [%v]: %s\n", m.name, level, message)
}

func main() {
	var logger Logger

	fLogger := NewFileLogger("FileLogger")
	logger = fLogger
	// 由于 FileLogger 提供了 doLog 的实现，这里会调用 FileLogger 的 doLog 方法
	logger.Log(LevelInfo, "This is an info message for file logger.")

	mqLogger := NewMessageQueueLogger("MQLogger")
	logger = mqLogger
	// 由于 MessageQueueLogger 提供了 doLog 的实现，这里会调用 MessageQueueLogger 的 doLog 方法
	logger.Log(LevelDebug, "This is a debug message for message queue logger.")
}
