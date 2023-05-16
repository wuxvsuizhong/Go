package fileop

import (
	"bufio"
	"fmt"
	"io"
	"mypro/getfuncinfo"
	"os"
)

/*
使用io.copy复制文件
io.copy 自带缓存，也可以拷贝大文件
*/
func FileCopy(dstFilename, srcFilename string) (n int64, err error) {
	srcfile, err := os.Open(srcFilename)
	if err != nil {
		fmt.Println("打开源文件失败!")
		return
	}
	defer srcfile.Close() //defer关闭打开的文件

	reader := bufio.NewReader(srcfile)

	dstfile, err := os.OpenFile(dstFilename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开目标文件失败!")
		return
	}
	defer dstfile.Close() //defer关闭打开的文件

	writer := bufio.NewWriter(dstfile)

	return io.Copy(writer, reader)
}

func TestCopyFile() {
	getfuncinfo.PrintFuncName()
	srcfile := "./a.jpg"
	dstfile := "./b.jpg"
	if _, err := FileCopy(dstfile, srcfile); err != nil {
		fmt.Println("拷贝文件失败! err=", err)
	} else {
		fmt.Println("拷贝完成!")
	}
}
