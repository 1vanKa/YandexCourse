package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	logger := NewLogExtended(os.Stderr, "Logger:")
	logger.SetLogLevel(LogLevelWarning)
	logger.Infoln("Не должно напечататься")
	logger.Warnln("Hello")
	logger.Errorln("World")
	logger.Println("Debug")
}

type LogLevel int

const (
	LogLevelError LogLevel = iota
	LogLevelWarning
	LogLevelInfo
)

type LogExtended struct {
	*log.Logger
	logLevel LogLevel
}

func NewLogExtended(out io.Writer, prefix string) (l *LogExtended) {
	l = &LogExtended{
		Logger:   log.New(out, prefix, log.LstdFlags),
		logLevel: LogLevelInfo,
	}
	return
}

func (l LogLevel) isValid() bool {
	switch l {
	case LogLevelError, LogLevelWarning, LogLevelInfo:
		return true
	default:
		return false
	}
}

func (l *LogExtended) SetLogLevel(lvl LogLevel) error {
	if !lvl.isValid() {
		return fmt.Errorf("not valid error: %d", lvl)
	}
	l.logLevel = lvl
	return nil
}

func (l *LogExtended) Infoln(msg string) {
	l.println(LogLevelInfo, "INFO ", msg)
}

func (l *LogExtended) Warnln(msg string) {
	l.println(LogLevelWarning, "WARN ", msg)
}

func (l *LogExtended) Errorln(msg string) {
	l.println(LogLevelError, "ERR ", msg)
}

func (l *LogExtended) println(logLvl LogLevel, prefix string, msg string) {
	if l.logLevel >= logLvl {
		l.Logger.Println(prefix + msg)
	}
}
