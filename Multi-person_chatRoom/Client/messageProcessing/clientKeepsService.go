package messageProcessing

import (
	"awesomeProject/Multi-person_chatRoom/Client/tools"
	"awesomeProject/Multi-person_chatRoom/Common/messageStruct"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func ShowRegisteredMenu(resMessageData *messageStruct.RegisteredResMessageData) {
	fmt.Printf("---------恭喜%v登录成功----------\n", resMessageData.UserName)
	fmt.Println("-------1.显示在线用户列表----------")
	fmt.Println(" ------2.发送消息-----------------")
	fmt.Println("-------3.信息列表-----------------")
	fmt.Println("-------4.退出系统-----------------")
	fmt.Println("       请选择(1 - 4):             ")
	var key int
	fmt.Scanln(&key)
	switch key {
	case 1:
		//临时显示
		fmt.Println("显示在线用户Id列表")
		for _, userId := range resMessageData.OnLineUsersId {
			//if userId == resMessageData.UserId {
			//	continue
			//}
			//初始化本地在线用户列表后台数据
			userStatus := &messageStruct.UserStatus{
				UserId:            userId,
				UserCurrentStatus: messageStruct.UserIsOnline,
			}
			onLineUsers[userId] = userStatus

			for userId_key, _ := range onLineUsers {
				if userId_key == resMessageData.UserId {
					continue
				}
				fmt.Printf("用户Id%v\n", userId_key)
			}
		}
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你选择退出了系统...")
		os.Exit(0)
	}

}

func ShowLoginMenu(resMessageData *messageStruct.LoginResMessageData) {
	fmt.Printf("\n\n---------恭喜%v登录成功----------\n", resMessageData.UserName)
	fmt.Println("-------1.显示在线用户列表----------")
	fmt.Println(" ------2.发送消息-----------------")
	fmt.Println("-------3.信息列表-----------------")
	fmt.Println("-------4.退出系统-----------------")
	fmt.Println("       请选择(1 - 4):             ")
	var key int
	fmt.Scanln(&key)
	var content string
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
		for _, userId := range resMessageData.OnLineUsersId {
			//if userId == resMessageData.UserId {
			//	continue
			//}
			//初始化本地在线用户列表后台数据
			userStatus := &messageStruct.UserStatus{
				UserId:            userId,
				UserCurrentStatus: messageStruct.UserIsOnline,
			}
			onLineUsers[userId] = userStatus
		}
		for userId_key, _ := range onLineUsers {
			if userId_key == resMessageData.UserId {
				continue
			}
			fmt.Printf("用户Id%v\n", userId_key)
		}
	case 2:
		fmt.Println("发送消息，要群发的消息")
		fmt.Scanln(&content)
		err := SsmsMessageProcessing{}.SendGroupMessage(content)
		if err != nil {
			fmt.Printf("群发短消息发送给服务器失败%v\n", err)
			return
		}
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你选择退出了系统...")
		os.Exit(0)
	}

}

func ClientKeepsService(coon net.Conn, userId int) {
	fmt.Println("客户端后台保持读取服务端连接的内容")
	tT := &tools.Transfer{
		Conn: coon,
	}
	for {
		mes, err := tT.ReedPKg()
		if err != nil {
			fmt.Printf("客户端保持读取服务端连接失败%v\n", err)
			return
		}
		switch mes.Type {

		case messageStruct.UserStatusNotificationType:
			//反序列化返回登录的消息
			userStatusNotification := messageStruct.UserStatusNotification{}
			err := json.Unmarshal([]byte(mes.Data), &userStatusNotification)
			if err != nil {
				fmt.Printf("当前登录用户反序列化失败%v\n", err)
				return
			}
			//-------------------存入客户端本地在线用户map列表
			//判断初始化时是否已经存在、避免重复录入
			userStatus, ok := onLineUsers[userStatusNotification.UserId]
			if !ok {
				userStatus = &messageStruct.UserStatus{
					UserId: userStatusNotification.UserId,
				}
			}
			userStatus.UserCurrentStatus = userStatusNotification.UserCurrentStatus

			onLineUsers[userStatusNotification.UserId] = userStatus

			for userId_key, _ := range onLineUsers {
				if userId_key == userId {
					continue
				}
				fmt.Printf("用户Id%v\n", userId_key)
			}
		case messageStruct.SmsResMessageType:
			fmt.Printf("接受到的群发消息结构体%v\n", mes)
			smsResMessageData := messageStruct.SmsResMessageData{}
			err := json.Unmarshal([]byte(mes.Data), &smsResMessageData)
			if err != nil {
				fmt.Printf("接受到的群发消息内容结构体", smsResMessageData)
			}
			fmt.Println(smsResMessageData.Content)
			fmt.Println(smsResMessageData.UserId)
		default:
			fmt.Println("客户端后台接收到服务端发来未知返回消息")
		}
	}
}
