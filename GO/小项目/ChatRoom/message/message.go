package message

const (
	LoginMsgType    = "LoginMsgType"
	RegisterMsgType = "RegisterMsgType"
	LoginType       = "LoginType"
	ResultMsgType   = "ResultMsgType"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginInfo struct {
	UserId   string `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type ResultMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
