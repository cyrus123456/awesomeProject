package main

import (
	"awesomeProject/Multi-person_chatRoom/Client/messageProcessing"
	"fmt"
	"os"
)

func main() {
	var key int

	var userId int
	userPwd := ""
	userName := ""

	for {
		fmt.Println("-----------------------欢迎登录多人聊天软件------------------------")
		fmt.Println("                          1、登录聊天软件                         ")
		fmt.Println("                          2、注册用户                            ")
		fmt.Println("                          3、退出系统                            ")
		fmt.Println("                          4、请选择（1-3）：                      ")

		fmt.Scanln(&key)

		switch key {
		case 1:
			fmt.Println("登录聊天软件")
			fmt.Println("请输如账号")
			fmt.Scanln(&userId)
			fmt.Println("请输密码")
			fmt.Scanln(&userPwd)
			messageProcessing := messageProcessing.UserMessageProcessing{}
			err := messageProcessing.Login(userId, userPwd)
			if err != nil {
				fmt.Printf("登陆失败%v\n", err)
			}
		case 2:
			fmt.Println("新用户注册")
			fmt.Println("请输如账号")
			fmt.Scanln(&userId)
			fmt.Println("请输密码")
			fmt.Scanln(&userPwd)
			fmt.Println("请输昵称")
			fmt.Scanln(&userName)
			messageProcessing := messageProcessing.UserMessageProcessing{}
			err := messageProcessing.Registered(userId, userPwd, userName)
			if err != nil {
				fmt.Printf("注册失败%v\n", err)
			}
		case 3:
			fmt.Println("退出登录")
			os.Exit(0)

		default:
			fmt.Println("请选择（1-3）:")
		}

	}
}
