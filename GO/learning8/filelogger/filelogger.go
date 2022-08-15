// filelogger
package filelogger

import (
	"errors"
	"fmt"
	"os"
	"path"
	"regexp"
	"runtime"
	"sort"
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

type FileLogger struct {
	level        LogLevel
	fileObj      *os.File
	errFileObj   *os.File
	filePath     string
	fileName     string
	fileSize     int64
	fileQuantity int64
}

func (f *FileLogger) levelMap(lvstr string) (LogLevel, error) {
	lvstr = strings.ToUpper(lvstr)
	switch lvstr {
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
		err := errors.New("未知的日志级别")
		return UNKNOWN, err
	}
}

func (f *FileLogger) getLogfile() (fp *os.File, errfp *os.File, err error) {
	//检查日志文件是否存在
	fpath := path.Join(f.filePath, f.fileName)
	fobj := f.fileObj
	finfo, err := os.Stat(fpath)
	if err == nil {
		//文件存在则检查大小
		if finfo.Size() >= f.fileSize {
			//分割文件
			err = f.fileObj.Close()
			fobj = nil
			if err != nil {
				fmt.Println("关闭文件失败,err:", err)
			} else {
				newfpath := fpath + time.Now().Format("20060102150405.000")
				os.Rename(fpath, newfpath)
			}
			f.removeFile()
		}
	}

	//若文件未打开则打开文件
	if fobj == nil {
		fobj, err = os.OpenFile(fpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("打开日志文件%s失败,err:%v\n", fpath, err)
			return nil, nil, err
		}
		f.fileObj = fobj //打开的日志文件描述符赋值给logger
	}

	errfpath := path.Join(f.filePath, f.fileName+".error")
	errfobj := f.errFileObj
	errfinfo, err := os.Stat(errfpath)
	if err == nil {
		if errfinfo.Size() >= f.fileSize {
			//重命名分割文件
			f.errFileObj.Close()
			errfobj = nil
			newerrfpath := errfpath + time.Now().Format("20060102150405.000")
			os.Rename(errfpath, newerrfpath)
			f.removeFile()
		}
	}
	//若文件未打开打开error日志文件
	if errfobj == nil {
		errfobj, err = os.OpenFile(errfpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("打开日志文件%s失败,err:%v\n", errfpath, err)
			return nil, nil, err
		}
		f.errFileObj = errfobj //打开的日志文件描述符赋值给logger
	}

	return fobj, errfobj, nil
}

func NewFileLogger(lvstr string, fpath string, fname string, size int64, quantity int64) (*FileLogger, error) {
	loghandler := &FileLogger{
		filePath:     fpath,
		fileName:     fname,
		fileSize:     size,
		fileQuantity: quantity,
	}

	level, err := loghandler.levelMap(lvstr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fobj, errfobj, err := loghandler.getLogfile()
	if err != nil {
		return nil, err
	}
	loghandler.level = level
	loghandler.fileObj = fobj
	loghandler.errFileObj = errfobj

	loghandler.removeFile() //触发一次老化检查

	return loghandler, nil

}

func (f *FileLogger) getlineinfo() (fileName, funcName string, lineno int) {
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

func (f *FileLogger) isenable(lv LogLevel) bool {
	return lv >= f.level
}

func (f *FileLogger) log(lvStr string, format string, a ...interface{}) {
	loglv, err := f.levelMap(lvStr)
	if err != nil {
		fmt.Printf("获取日志级别映射失败,err:%v\n", err)
		return
	}

	if !f.isenable(loglv) {
		//日志级别低于初始化时候设置的日志级别时不输出日志
		return
	}

	_, _, err = f.getLogfile()
	if err != nil {
		return
	}
	msg := fmt.Sprintf(format, a...)
	now := time.Now().Format("2006-01-01 15:04:05")
	file, funcname, lineno := f.getlineinfo()
	fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d]-- %s\n", now, lvStr, file, funcname, lineno, msg)
	if loglv >= ERROR {
		//记录error及error级别以上的日志到error日志文件中
		fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d]-- %s\n", now, lvStr, file, funcname, lineno, msg)
	}
}

func (f *FileLogger) getfilelist() (logArr []string, errArr []string) {
	var logfileArr []string
	var errfileArr []string
	re, _ := regexp.Compile(`^` + f.fileName + `\d+`)
	errRe, _ := regexp.Compile(`^` + f.fileName + ".error" + `\d+`)
	flist, _ := os.ReadDir(path.Join(f.filePath))
	for _, item := range flist {
		itemName := path.Base(item.Name())
		if re.Match([]byte(itemName)) {
			//log 日志
			logfileArr = append(logfileArr, itemName)
		} else if errRe.Match([]byte(itemName)) {
			//error 日志
			errfileArr = append(errfileArr, itemName)
		}
	}
	return logfileArr, errfileArr
}

func (f *FileLogger) removeFile() {
	loglist, errlist := f.getfilelist()
	for _, list := range [][]string{loglist, errlist} {
		increment := len(list) - int(f.fileQuantity)
		if increment > 0 {
			sort.Slice(list, func(i, j int) bool {
				return list[i] < list[j]
			})
			for i := 0; i < increment; i++ {
				err := os.Remove(path.Join(f.filePath, list[i]))
				if err != nil {
					fmt.Printf("老化历史文件出错,err:%v\n", err)
				}
			}
		}
	}
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log("INFO", format, a...)
}

func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log("DEBUG", format, a...)
}

func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log("WARNING", format, a...)
}

func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log("ERROR", format, a...)
}

func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log("FATAL", format, a...)
}
