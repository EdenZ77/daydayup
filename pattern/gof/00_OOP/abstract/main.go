package main

import (
	"fmt"
	"time"
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

// Log 这是一个模板方法，它定义了日志记录的通用步骤
func (l *BaseLogger) Log(level Level, message string) {
	// 通用操作：添加时间戳
	timestamp := time.Now().Format(time.RFC3339)
	formattedMessage := fmt.Sprintf("[%s] [%s]: %s", timestamp, l.name, message)

	// 调用 doLog 方法，由子类实现具体的日志记录方式
	l.doLog(level, formattedMessage)
}

// doLog 抽象方法，子类必须提供实现
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
	fmt.Printf("[File] %s\n", message)
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
	fmt.Printf("[MessageQueue] %s\n", message)
}

func main() {
	var logger Logger

	fLogger := NewFileLogger("FileLogger")
	logger = fLogger
	// 使用模板方法 Log，它包含了通用的日志记录步骤，以及特定于 FileLogger 的步骤
	logger.Log(LevelInfo, "This is an info message for file logger.")

	mqLogger := NewMessageQueueLogger("MQLogger")
	logger = mqLogger
	// 使用模板方法 Log，它包含了通用的日志记录步骤，以及特定于 MessageQueueLogger 的步骤
	logger.Log(LevelDebug, "This is a debug message for message queue logger.")
}
