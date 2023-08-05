package message

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func ReadPack(conn net.Conn, msg *Message) (err error) {
	var buf [8192]byte
	_, err = conn.Read(buf[:4])
	if err != nil {
		if err == io.EOF {
			fmt.Println("客户端断开连接,服务端关闭coekct!")
			return
		}
		fmt.Println("conn.Read msgLen err:", err)
		return
	}
	msgLen := binary.BigEndian.Uint32(buf[:4])
	fmt.Println("接收到数据长度:", msgLen)
	n, err := conn.Read(buf[:msgLen])
	if err != nil || uint32(n) != msgLen {
		fmt.Println("conn.Read data err:", err)
		return
	}
	err = json.Unmarshal(buf[:msgLen], msg)
	if err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}
	return
}

func SendMsg(conn net.Conn, msg *Message) (err error) {
	msgBytes, err := json.Marshal(*msg)
	if err != nil {
		fmt.Println("writeMsg json.Marshal err:", err)
		return
	}

	//组装发送的消息数据(消息长度+消息数据)
	var lenBytes [4]byte
	binary.BigEndian.PutUint32(lenBytes[:4], uint32(len(msgBytes)))
	var wholeMsgBytes []byte
	wholeMsgBytes = append(wholeMsgBytes, lenBytes[:]...)
	wholeMsgBytes = append(wholeMsgBytes, msgBytes...)

	n, err := conn.Write(wholeMsgBytes)
	if err != nil || n != len(wholeMsgBytes) {
		fmt.Println("writeMsg conn.Write err:", err)
	}

	return
}