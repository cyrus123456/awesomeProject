package messageStruct

const (
	LoginMessageType                   = "loginMessageType"
	LoginResMessageType                = "loginResMessageType"
	RegisteredMessageType              = "registeredMessageType"
	RegisteredResMesRegisteredsageType = "registeredResMessageType"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMessageData struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMessageData struct {
	Code     int    `json:"code"` //返回的状态码
	Error    string `json:"error"`
	UserName string `json:"userName"`
}

type RegisteredMessageData struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type RegisteredResMessageData struct {
	Code     int    `json:"code"` //返回的状态码
	Error    string `json:"error"`
	UserName string `json:"userName"`
}
