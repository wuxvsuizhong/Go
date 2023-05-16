package fileop

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func TestFuncWrite() { //
	//文件名,创建方式,文件权限(权限控制只在linux下生效)
	//fileobj,err := os.OpenFile("./xxx.txt",os.O_WRONLY|os.O_CREATE|os.O_APPEND,0644)
	fileobj, err := os.OpenFile("./xxx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	//O_TRUNC 每次打开都重新写入
	if err != nil {
		fmt.Printf("打开文件失败!")
		return
	}
	defer fileobj.Close()

	fileobj.Write([]byte("测试文件文件写入byte。。。\n"))
	fileobj.WriteString("测试直接写入string。。。\n")

}

func TestBufioWr() {
	fileobj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("打开文件失败!")
		return
	}
	defer fileobj.Close()
	wr := bufio.NewWriter(fileobj)              //bufio.Writer是自带缓冲的，所以Flush操作时必要的
	wr.WriteString("测试写入string,use bufio。。。\n") //把内容写到bufio缓存中
	wr.Flush()                                  //把缓存内容写入文件中
}

func TestIoutilWR() {
	str := "直接使用ioutil写入byte。。。\n"
	err := ioutil.WriteFile("./xxxxx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("ioutil写文件失败")
		return
	}
}
