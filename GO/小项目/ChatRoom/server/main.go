package main

import (
  "chatroom/message"
  "encoding/binary"
  "encoding/json"
  "errors"
  "fmt"
  "io"
  "net"
)

func main() {
  listen, err := net.Listen("tcp", "0.0.0.0:8889")
  if err != nil {
    fmt.Println("net.Listen err:", err)
    return
  }
  defer listen.Close()

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
    fmt.Println("捕获到异常:", exception)
    return exception.(error)
  }()
  defer conn.Close()
  for {
    var msg message.Message
    err := readMsg(conn, &msg)
    if err != nil {
      if err == io.EOF {
        fmt.Println("客户端断开连接,服务端关闭coekct!")
        return
      }
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

func readMsg(conn net.Conn, msg *message.Message) (err error) {
  var buf [8192]byte
  _, err = conn.Read(buf[:4])
  if err != nil {
    fmt.Println("conn.Read msgLen err:", err)
    return
  }
  msgLen := binary.BigEndian.Uint32(buf[:4])
  fmt.Println("接受到数据长度:", msgLen)
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

func writeMsg(conn net.Conn, msg *message.Message) (err error) {
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

func serverProcessMsg(conn net.Conn, msg *message.Message) (err error) {

  switch msg.Type {
  case message.LoginType:
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
  if info.UserId != "100" {
    //fmt.Println("登录失败!")
    panic(errors.New("登录失败"))
  }

  var resinfo message.ResultMsg
  resinfo.Code = 200
  resinfo.Msg = "登录成功!"

  var resMsg message.Message
  resMsg.Type = message.ResultMsgType
  dataBytes, err := json.Marshal(resinfo)
  if err != nil {
    fmt.Println("serverProcessLogin json.Marshal err:", err)
    return
  }
  resMsg.Data = string(dataBytes)
  if err = writeMsg(conn, &resMsg); err != nil {
    fmt.Println("serverProcessLogin writeMsg err:", err)
    return
  }

  return
}
