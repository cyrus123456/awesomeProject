package messageProcessing

import "awesomeProject/Multi-person_chatRoom/Common/messageStruct"

var onLineUsers map[int]*messageStruct.UserStatus = make(map[int]*messageStruct.UserStatus, 1024)
