package login

import (
	"chatroom/message"
	"encoding/json"
	"fmt"
	"net"
)

func Login() (err error) {
	conn, err := net.Dial("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()

	var msg message.Message
	msg.Type = message.LoginMsgType

	var loginInfo message.LoginInfo
	fmt.Print("请输入用户ID:")
	fmt.Scanf("%s\n", &loginInfo.UserId)
	fmt.Print("请输入密码:")
	fmt.Scanf("%s\n", &loginInfo.UserPwd)
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
		fmt.Println(retMsg.Msg)
	} else {
		fmt.Println("数据错误!,err:")
	}

	return
}
