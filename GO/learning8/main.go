package main

import (
	"logger/filelogger"
	"logger/logger"
	"time"
)

func testConsoleLog() {
	log, err := logger.NewLog("info")
	if err != nil {
		return
	}
	log.Info("info 字符串")
	log.Debug("debug 字符串")
	log.Warning("warning 字符串")
	log.Error("error 字符串")
	log.Fatal("fatal 字符串")
	key := 10086
	val := "中国移动"
	log.Info("info 消息日志 %d:%s", key, val)
}

func testFilelog() {
	log, err := filelogger.NewFileLogger("info", "./", "logtest.log", 10*1024, 5)
	if err != nil {
		return
	}
	for {
		log.Info("info 字符串")
		log.Debug("debug 字符串")
		log.Warning("warning 字符串")
		log.Error("error 字符串")
		log.Fatal("fatal 字符串")
		key := 10086
		val := "中国移动"
		log.Info("info 消息日志 %d:%s", key, val)
		time.Sleep(200 * time.Millisecond)
	}

}

func main() {
	// testConsoleLog()
	testFilelog()
}
