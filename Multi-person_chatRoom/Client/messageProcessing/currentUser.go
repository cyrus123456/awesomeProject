package messageProcessing

import (
	"awesomeProject/Multi-person_chatRoom/Common/messageStruct"
	"net"
)

type CurrentUser struct {
	Coon net.Conn
	messageStruct.UserStatus
}

var currentUser CurrentUser
