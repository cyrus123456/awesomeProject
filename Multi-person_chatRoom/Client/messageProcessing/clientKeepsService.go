package messageProcessing

import (
	"awesomeProject/Multi-person_chatRoom/Client/tools"
	"fmt"
	"net"
	"os"
)

func ShowLoginMenu(registeredUserName string) {
	fmt.Printf("---------恭喜%v登录成功----------\n", registeredUserName)
	fmt.Println("-------1.显示在线用户列表----------")
	fmt.Println(" ------2.发送消息-----------------")
	fmt.Println("-------3.信息列表-----------------")
	fmt.Println("-------4.退出系统-----------------")
	fmt.Println("       请选择(1 - 4):             ")
	var key int
	fmt.Scanln(&key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你选择退出了系统...")
		os.Exit(0)
	}

}

func ClientKeepsService(coon net.Conn) {
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
		fmt.Printf("客户端保持读取服务端连接的内容%v\n", mes)
	}
}
