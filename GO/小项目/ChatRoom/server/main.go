package main

import (
	"chatroom/message"
	"chatroom/redisop"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listen.Close()
	// 初始化redis连接池
	redisop.InitPool("127.0.0.1:6379", 5, 20, 300*time.Second)

	for {
		fmt.Println("服务器等待客户端连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err:", err)
			return
		}
		go process(conn)
	}

}

func process(conn net.Conn) {
	defer func() (err error) {
		exception := recover()
		//errText := fmt.Sprintf("serverProcessMsg exception,err:%s", exception)
		//err = errors.New(errText)
		if exception != nil {
			fmt.Println("server捕获到异常:", exception)
		}
		return
	}()
	defer conn.Close()
	for {
		var msg message.Message
		//err := readMsg(conn, &msg)
		err := msg.ReadPack(conn)
		if err != nil {
			fmt.Println("readMsg err:", err)
			return
		}
		fmt.Println("服务器接受到客户端消息:", msg)
		err = serverProcessMsg(conn, &msg)
		if err != nil {
			fmt.Println(err)
		}

	}

}

func serverProcessMsg(conn net.Conn, msg *message.Message) (err error) {

	switch msg.Type {
	case message.LoginMsgType:
		fmt.Println("服务端去处理登录!")
		serverProcessLogin(conn, msg.Data)
	case message.RegisterMsgType:
		fmt.Println("服务端去处理注册!")
	default:
		fmt.Println("未知的消息类型!")
		return
	}

	return
}

func serverProcessLogin(conn net.Conn, data string) (err error) {
	var info message.LoginInfo
	err = json.Unmarshal([]byte(data), &info)
	if err != nil {
		fmt.Println("serverProcessLogin json.Unmarshal err:", err)
		return
	}
	fmt.Println(info)

	var resinfo message.ResultMsg
	_, err = redisop.QueryHashKeyVal("user", info.UserId)
	if err != nil {
		resinfo.Code = 400
		resinfo.Msg = "没找到用户信息，请先注册!"
	} else {
		resinfo.Code = 200
		resinfo.Msg = "登录成功!"
	}

	var resMsg message.Message
	resMsg.Type = message.LoginMsgRes
	dataBytes, err := json.Marshal(resinfo)
	if err != nil {
		fmt.Println("serverProcessLogin json.Marshal err:", err)
		return
	}
	resMsg.Data = string(dataBytes)
	//返回客户端的登录结果消息
	if err = resMsg.SendPack(conn); err != nil {
		fmt.Println("serverProcessLogin writeMsg err:", err)
		return
	}

	return
}
