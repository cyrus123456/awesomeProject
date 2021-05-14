package messageProcessing

import (
	"awesomeProject/Multi-person_chatRoom/Common/messageStruct"
	"awesomeProject/Multi-person_chatRoom/Server/tools"
	"encoding/json"
	"fmt"
	"net"
)

type SmsMessageProcessing struct{}

func (this SmsMessageProcessing) Group_messageProcessing(mes *messageStruct.Message) (err error) {
	//实例化一个空的注册消息结构体
	smsMessageData := messageStruct.SmsMessageData{}

	//把服务端接收到的序列化的登录消息反序列化然后存到&registeredMesData结构体,看一眼
	err = json.Unmarshal([]byte(mes.Data), &smsMessageData)
	if err != nil {
		fmt.Printf("服务器处理服务端群发消息反序列化存入群发消息结构体失败%v\n", err)
		return
	}
	fmt.Printf("服务端群发消息内容%v\n", smsMessageData)
	mes.Type = messageStruct.SmsResMessageType

	//看完内容直接序列化再返回
	mesByte, err := json.Marshal(mes)
	if err != nil {
		fmt.Printf("服务端群发消息结构体序列化失败%v\n", err)
		return
	}
	//开始群发消息
	for key, userMessageProcessing := range onLineUserMgr.onLineUsers {
		if key == smsMessageData.UserId {
			continue
		}
		err := userMessageProcessing.GroupPosting(mesByte, userMessageProcessing.Coon)
		if err != nil {
			fmt.Printf("服务端群发消息函数失败%v\n", err)
		}
	}
	return
}

//群发函数函数
func (this *UserMessageProcessing) GroupPosting(mesByte []byte, coon net.Conn) (err error) {
	//	发送返回客户端
	tT := tools.Transfer{
		Conn: coon,
	}
	mes, err := tT.WritePKg(mesByte)
	if err != nil {
		fmt.Printf("服务端群发消息给其他用户失败%v\n", err)
		return
	}
	fmt.Printf("服务端群发消息给其他用户的信息%v\n", mes)
	return
}
