package clinetrequest

import (
	"chatroom/message"
	"encoding/json"
	"fmt"
	"net"
)

type connInfo struct {
	conn net.Conn
}

var (
	conninfo = connInfo{}
)

func connct2Server() (conn net.Conn, err error) {
	if conninfo.conn != nil {
		return conninfo.conn, nil
	}
	conn, err = net.Dial("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	conninfo.conn = conn
	// defer conn.Close()
	return
}

func getUserInfo() (loginInfo message.LoginInfo) {
	fmt.Print("请输入用户ID:")
	fmt.Scanf("%s\n", &loginInfo.UserId)
	fmt.Print("请输入密码:")
	fmt.Scanf("%s\n", &loginInfo.UserPwd)
	return
}

func errorProcess() {
	exception := recover()
	if exception != nil {
		fmt.Println("ErrorProcess 捕获到异常:", exception)
		conninfo.conn.Close()
		conninfo.conn = nil
		return
	}
}

func Login() (err error) {
	defer errorProcess()

	// 连接服务器
	conn, _ := connct2Server()

	var msg message.Message
	msg.Type = message.LoginMsgType

	var loginInfo = getUserInfo()
	data, err := json.Marshal(loginInfo)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
	}
	msg.Data = string(data)

	if err = msg.SendPack(conn); err != nil {
		fmt.Println("SendMsg err:", err)
		return
	}
	//读取响应结果
	var retPack message.Message
	err = retPack.ReadPack(conn)
	if err != nil {
		fmt.Println("Login message.ReadMsg err:", err)
	}
	fmt.Println("收到的服务器的返回pack:", retPack)
	if retPack.Type == message.LoginMsgRes {
		var retMsg message.ResultMsg
		if err = json.Unmarshal([]byte(retPack.Data), &retMsg); err != nil {
			fmt.Println("json.Unmarshal err:", err)
		}
		fmt.Println("收到的服务器的返回msg:", retMsg)
		if retMsg.Code != 200 {
			fmt.Println("登录失败!")
			return
		}
		fmt.Println(retMsg.Msg) //登录成功
	} else {
		fmt.Println("数据错误!,err:")
	}

	return
}

func Registery() (err error) {
	defer errorProcess()
	conn, err := connct2Server()

	loginInfo := getUserInfo()
	data, err := json.Marshal(loginInfo)
	var msg message.Message
	msg.Type = message.RegisterMsgType
	msg.Data = string(data)
	msg.SendPack(conn)

	var retMsg message.Message
	if err = retMsg.ReadPack(conn); err != nil {
		fmt.Println("Registery ReadPack err:", err)
		
	}

	return
}
