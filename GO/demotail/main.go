package main

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

func main() {
	fName := "./my.log"
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}

	tails, err := tail.TailFile(fName, config)
	if err != nil {
		fmt.Printf("TailFile 出错,err:%v\n", err)
		return
	}

	var (
		line *tail.Line
		ok   bool
	)

	for {
		line, ok = <-tails.Lines
		if !ok {
			fmt.Printf("跟踪文件已关闭:%s,尝试重连...\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println(line.Text)
	}
}
