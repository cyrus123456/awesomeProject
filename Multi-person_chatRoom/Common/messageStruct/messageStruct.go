package messageStruct

//消息类型敞亮
const (
	LoginMessageType                   = "loginMessageType"
	LoginResMessageType                = "loginResMessageType"
	RegisteredMessageType              = "registeredMessageType"
	RegisteredResMesRegisteredsageType = "registeredResMessageType"
	UserStatusNotificationType         = "userStatusNotificationType"
	SmsMessageType                     = "smsMessageType"
	SmsResMessageType                  = "smsResMessageType"
)

//用户状态常量
const (
	UserIsOnline  = "userIsOnline"
	userIsOffline = "userIsOffline"
	userIsBusy    = "userIsBusy"
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
	OnLineUsersId []int  `json:"onLineUsersId"`
	UserId        int    `json:"userId"`
	Code          int    `json:"code"` //返回的状态码
	Error         string `json:"error"`
	UserName      string `json:"userName"`
}

type RegisteredMessageData struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type RegisteredResMessageData struct {
	OnLineUsersId []int  `json:"onLineUsersId"`
	UserId        int    `json:"userId"`
	Code          int    `json:"code"` //返回的状态码
	Error         string `json:"error"`
	UserName      string `json:"userName"`
}

type UserStatusNotification struct {
	UserId            int    `json:"userId"`
	UserCurrentStatus string `json:"userCurrentStatus"`
}

type UserStatus struct {
	UserId            int    `json:"userId"`
	UserPwd           string `json:"userPwd"`
	UserName          string `json:"userName"`
	UserCurrentStatus string `json:"userCurrentStatus"`
}

type SmsMessageData struct {
	Content    string `json:"content"`
	UserStatus `json:"userStatus"`
}

type SmsResMessageData struct {
	Content    string `json:"content"`
	UserStatus `json:"userStatus"`
}
