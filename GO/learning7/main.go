package main

import(
	"logger/mylogger"
)

func main(){
	log := mylogger.NewLog("debug")
	log.Debug("Debug字符串")
	log.Info("info 字符串")
	log.Error("Error 字符串")
	log.Fatal("Fatal 字符串")
	log.Warning("Warning 字符串")
}
