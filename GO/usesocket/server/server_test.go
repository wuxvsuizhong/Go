package server

import (
	"fmt"
	"net"
	"testing"
)

func process(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("客户端已退出:", err)
			return
		}
		fmt.Print(string(buf[:n])) //需要切片取读到的长度，否则会打印整个的buf
	}
}

func TestServerStart(t *testing.T) {
	fmt.Println("服务为启动...")
	listen, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		fmt.Println("listenm err:", err)
	}
	defer listen.Close()

	for {
		fmt.Println("等待客户端来连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept err:", err)
		} else {
			fmt.Println("收到客户端连接！客户端IP:", conn.RemoteAddr().String())
			go process(conn)
		}

	}
}
