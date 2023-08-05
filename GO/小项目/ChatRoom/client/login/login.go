package login

import (
	"chatroom/message"
	"encoding/binary"
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

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marchal err:", err)
	}
	msgLen := uint32(len(msgBytes))
	var lenBytes = make([]byte, 4)
	binary.BigEndian.PutUint32(lenBytes, msgLen)

	var wholeBytes []byte
	wholeBytes = append(wholeBytes, lenBytes...)
	wholeBytes = append(wholeBytes, msgBytes...)

	n, err := conn.Write(wholeBytes)
	if err != nil || n != len(wholeBytes) {
		fmt.Println("conn.Write err:", err)
	}
	//读取响应结果
	var retPack message.Message
	err = message.ReadPack(conn, &retPack)
	if err != nil {
		fmt.Println("Login message.ReadMsg err:", err)
	}
	if retPack.Type == message.LoginMsgRes {
		var retMsg message.ResultMsg
		json.Unmarshal([]byte(retPack.Data), retMsg)
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
