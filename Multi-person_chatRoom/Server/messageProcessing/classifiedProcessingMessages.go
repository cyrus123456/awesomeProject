package messageProcessing

import (
	"awesomeProject/Multi-person_chatRoom/Common/messageStruct"
	"awesomeProject/Multi-person_chatRoom/Server/tools"
	"fmt"
	"io"
	"net"
)

type Interface_ClassifiedProcessing_Messages interface {
	ClassifiedProcessingMessages(mes *messageStruct.Message) error
	CommunicationFor() error
}

type ClassifiedProcessing_Messages struct {
	Coon net.Conn
}

//根据不同消息类型处理不同消息
func (this *ClassifiedProcessing_Messages) ClassifiedProcessingMessages(mes *messageStruct.Message) (err error) {
	for {
		switch mes.Type {

		case messageStruct.LoginMessageType:
			fmt.Println("处理登陆消息")
			userMessageProcessing := &UserMessageProcessing{
				Coon: this.Coon,
			}
			err = userMessageProcessing.Login_messageProcessing(mes)
			if err != nil {
				fmt.Printf("处理登陆消息失败%v\n", err)
				return
			}
		case messageStruct.RegisteredMessageType:
			fmt.Println("处理注册消息")
			userMessageProcessing := &UserMessageProcessing{
				Coon: this.Coon,
			}
			err = userMessageProcessing.Registered_messageProcessing(mes)
			if err != nil {
				fmt.Printf("处理注册消息失败%v\n", err)
				return
			}
		case messageStruct.SmsMessageType:
			fmt.Printf("客户端收到群发消息%v\n", mes)
			smsMessageProcessing := SmsMessageProcessing{}
			smsMessageProcessing.Group_messageProcessing(mes)
		default:
			fmt.Println("消息类型错误")
		}
		return
	}

}

func (this *ClassifiedProcessing_Messages) CommunicationFor() (Err error) {

	for {
		tF := tools.Transfer{
			Conn: this.Coon,
		}
		mes, err := tF.ReedPKg()
		if err != nil {
			if err == io.EOF {
				fmt.Printf("客户端退出%v\n", err)
				//Err = err
				return
			} else {
				fmt.Printf("服务端读取失败%v\n", err)
				//Err = err
				return
			}
		}
		fmt.Printf("服务器对客户端发送来的信息(反序列化以后)%v\n", mes)

		err = this.ClassifiedProcessingMessages(&mes)
		if err != nil {
			fmt.Printf("服务器信息结构体分类处理失败%v\n", err)
			//Err = err
			return
		}
	}

}
