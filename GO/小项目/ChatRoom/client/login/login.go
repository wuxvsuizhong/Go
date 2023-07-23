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
	msg.Type = message.LoginType

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
	return
}
