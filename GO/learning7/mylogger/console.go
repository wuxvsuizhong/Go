package mylogger

import(
	"fmt"
	"time"
	"strings"
	"errors"
	"runtime"
	"path"
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

func parseloglever(s string)(LogLevel,error) {
	s = strings.ToLower(s)
	switch s{
	case "debug":
		return DEBUG,nil
	case "info":
		return INFO,nil
	case "warning":
		return WARNING,nil
	case "error":
		return ERROR,nil
	case "fatal":
		return FATAL,nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN,err
	}
}


type Logger struct{
	level LogLevel
}

func NewLog(levelStr string) Logger{
	level,err := parseloglever(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{level:level,}
}

func (l Logger)enable(loglevel LogLevel)bool{
	return loglevel >= l.level
}

func log(lv string,msg string){
	now := time.Now()
	fileName,funcName,lineno := getlineinfo()
	lv = strings.ToUpper(lv)
	fmt.Printf("[%s] [%s] [%s:%s:%d]-- %s\n",now.Format("2006-01-02 15:04:05"),lv,fileName,funcName,lineno,msg)
}

func (l Logger)Debug(s string){
	if l.enable(DEBUG) {
		//now := time.Now()
		//fileName,funcName,lineno := getlineinfo()
		//fmt.Printf("[%s] [DEBUG] [%s:%s:%d]-- %s\n",now.Format("2006-01-02 15:04:05"),fileName,funcName,lineno,s)
		log("debug",s)
	}
}

func (l Logger)Info(s string){
	if l.enable(INFO){
		//now := time.Now()
		//fileName,funcName,lineno := getlineinfo()
		//fmt.Printf("[%s] [INFO] [%s:%s:%d]-- %s\n",now.Format("2006-01-02 15:04:05"),fileName,funcName,lineno,s)
		log("info",s)
	}
}

func (l Logger)Error(s string){
	if l.enable(ERROR){
		//now := time.Now()
		//fileName,funcName,lineno := getlineinfo()
		//fmt.Printf("[%s] [ERROR] [%s,%s,%d]-- %s\n",now.Format("2006-01-02 15:04:05"),fileName,funcName,lineno,s)
		log("error",s)
	}
}

func (l Logger)Warning(s string){
	if l.enable(WARNING){
		//now := time.Now()
		//fileName,funcName,lineno := getlineinfo()
		//fmt.Printf("[%s] [WARNING] [%s,%s,%d]-- %s\n",now.Format("2006-01-02 15:04:05"),fileName,funcName,lineno,s)
		log("warning",s)
	}
}

func (l Logger)Fatal(s string){
	if l.enable(FATAL){
		//now := time.Now()
		//fileName,funcName,lineno := getlineinfo()
		//fmt.Printf("[%s] [FATAL] [%s,%s,%d]-- %s\n",now.Format("2006-01-02 15:04:05"),fileName,funcName,lineno,s)
		log("fatal",s)
	}
}

func getlineinfo()(funcName,fileName string,line int){
	pc,filepath,line,ok := runtime.Caller(3)	//栈向上三层找到主调函数
	if !ok {
		fmt.Printf("getlineinfo error\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()		//获取函数名称
	funcName = strings.Split(funcName,".")[1]	//获取package.func 的后半段方法名
	fileName = path.Base(filepath)
	return fileName,funcName,line
}
