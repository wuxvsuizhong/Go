package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func TestStartClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("client dial err:", err)
		return
	}
	defer conn.Close()
	fmt.Println("conn 成功!", conn)

	reader := bufio.NewReader(os.Stdin) //reader读取终端的输入
	for {
		fmt.Print(">>")
		data, err := reader.ReadString('\n') //获取reader读取的数据
		if err != nil {
			fmt.Println("从终端读取数据失败!err:", err)
		}
		if strings.Trim(data, "\r\n") == "exit" { //客户端输入exit的时候退出
			fmt.Println("客户端退出了!")
			break
		}

		_, err = conn.Write([]byte(data)) //conn写数据，也就是发送数据
		if err != nil {
			fmt.Println("conn.Write err=", err)
		}
	}

	//fmt.Printf("客户端发送了%v个字节后，退出!\n", n)
}
