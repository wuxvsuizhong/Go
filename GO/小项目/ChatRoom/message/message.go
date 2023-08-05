package message

type MsgType uint64

const (
	LoginMsgType MsgType = iota
	RegisterMsgType
	LoginMsgRes
	ResultMsgRes
)

type Message struct {
	Type MsgType `json:"type"`
	Data string  `json:"data"` //具体的message结构体
}

type LoginInfo struct {
	UserId   string `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type ResultMsg struct {
	Code int    `json:"code"` //返回状态码
	Msg  string `json:"msg"`  //消息体
}
