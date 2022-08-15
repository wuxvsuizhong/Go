package logger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

type LogLevel uint16

const (
	UNKNOWN LogLevel = iota
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

type Logger struct {
	level LogLevel
}

func (l Logger) levelMap(lvStr string) (LogLevel, error) {
	lvStr = strings.ToUpper(lvStr)
	switch lvStr {
	case "DEBUG":
		return DEBUG, nil
	case "INFO":
		return INFO, nil
	case "WARNING":
		return WARNING, nil
	case "ERROR":
		return ERROR, nil
	case "FATAL":
		return FATAL, nil
	default:
		err := errors.New("未知的日志级别!")
		return UNKNOWN, err
	}
}

func NewLog(lvStr string) (Logger, error) {
	level, err := Logger{}.levelMap(lvStr)
	if err != nil {
		fmt.Println("初始化Logger失败!")
		return Logger{}, err
	}
	return Logger{level: level}, nil
}

func (l Logger) getlineinfo() (fileName, funcName string, lineno int) {
	pc, fileName, lineno, ok := runtime.Caller(3)
	if !ok {
		fmt.Println("runtime caller 调用失败!")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	funcName = strings.Split(funcName, ".")[1]
	fileName = path.Base(fileName)
	return fileName, funcName, lineno
}

func (l Logger) log(lvStr string, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now().Format("2006-01-01 15:04:05")
	file, funcname, lineno := l.getlineinfo()
	fmt.Printf("[%s] [%s] [%s:%s:%d]-- %s\n", now, lvStr, file, funcname, lineno, msg)
}

func (l Logger) isenable(lvStr string) bool {
	level, err := l.levelMap(lvStr)
	if err != nil {
		panic(err)
	}
	if level >= l.level {
		return true
	}
	return false
}

func (l Logger) Info(format string, a ...interface{}) {
	if l.isenable("INFO") {
		l.log("INFO", format, a...)
	}
}

func (l Logger) Debug(format string, a ...interface{}) {
	if l.isenable("DEBUG") {
		l.log("DEBUG", format, a...)
	}
}

func (l Logger) Warning(format string, a ...interface{}) {
	if l.isenable("WARNING") {
		l.log("WARNING", format, a...)
	}
}

func (l Logger) Error(format string, a ...interface{}) {
	if l.isenable("ERROR") {
		l.log("ERROR", format, a...)
	}
}

func (l Logger) Fatal(format string, a ...interface{}) {
	if l.isenable("FATAL") {
		l.log("FATAL", format, a...)
	}
}
