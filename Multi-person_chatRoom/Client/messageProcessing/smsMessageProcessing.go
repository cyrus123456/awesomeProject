package messageProcessing

import (
	"awesomeProject/Multi-person_chatRoom/Client/tools"
	"awesomeProject/Multi-person_chatRoom/Common/messageStruct"
	"encoding/json"
	"fmt"
)

type SsmsMessageProcessing struct{}

func (this SsmsMessageProcessing) SendGroupMessage(content string) (err error) {

	smsMessageData := messageStruct.SmsMessageData{
		Content: content,
	}
	smsMessageData.UserId = currentUser.UserStatus.UserId
	smsMessageData.UserName = currentUser.UserStatus.UserName
	smsMessageData.UserCurrentStatus = currentUser.UserStatus.UserCurrentStatus
	smsMessageData_json, err := json.Marshal(smsMessageData)
	if err != nil {
		fmt.Printf("短消息内容结构体序列化失败%v\n", err)
		return
	}
	mes := messageStruct.Message{
		Type: messageStruct.SmsMessageType,
		Data: string(smsMessageData_json),
	}
	mes_jsonByte, err := json.Marshal(mes)
	if err != nil {
		fmt.Printf("群发消息发送结构体序列化失败%v\n", err)
		return
	}
	tT := &tools.Transfer{
		Conn: currentUser.Coon,
	}
	mes, err = tT.WritePKg(mes_jsonByte)
	if err != nil {
		fmt.Printf("群发短消息发送给服务器失败%v\n", err)
		return
	}

	return
}
