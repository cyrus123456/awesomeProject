package messageProcessing

import (
	"awesomeProject/Multi-person_chatRoom/Client/tools"
	"awesomeProject/Multi-person_chatRoom/Common/messageStruct"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

type UserMessageProcessing struct{}

func (this *UserMessageProcessing) Registered(userId int, userPwd, userName string) (err error) {
	fmt.Printf("登录输入的账号ID%v密码%v昵称%v\n", userId, userPwd, userName)
	//拨号连接
	coon, err := net.Dial("tcp", "0.0.0.0:6666")
	if err != nil {
		fmt.Printf("\"0.0.0.0:6666\"拨号失败%v\n", err)
		return
	}
	//释放连接资源
	defer coon.Close()
	registeredData := messageStruct.RegisteredMessageData{
		userId,
		userPwd,
		userName,
	}

	//注册信息序列化

	registeredData_json, err := json.Marshal(registeredData)
	if err != nil {
		fmt.Printf("注册消息序列化失败%v\n", err)
		return
	}

	//注册信息放入添加到完整消息结构体
	mes := messageStruct.Message{
		messageStruct.RegisteredMessageType,
		string(registeredData_json),
	}
	//装有内容的消息结构体序列化
	mes_jsonByte, err := json.Marshal(mes)
	if err != nil {
		fmt.Printf("准备发送诸注册消息结构体序列化为字节切片失败%v\n", err)
		return
	}

	//-------------------------准备发送消息长度

	tT := tools.Transfer{
		Conn: coon,
	}

	mes, err = tT.WritePKg(mes_jsonByte)

	if err != nil {
		fmt.Printf("注册消息写入发送失败%v\n", err)
		return
	}

	//-------------------读取客户端返回的消息

	reedPKg_mes, err := tT.ReedPKg()
	if err != nil {
		if err == io.EOF {
			fmt.Printf("客户端退出%v\n", err)
			return
		} else {
			fmt.Printf("服务端读取失败%v\n", err)
			return
		}
	}
	fmt.Printf("客户端接收读取服务器发送来的信息(结构体)%v\n", reedPKg_mes)
	registeredResMessageData := messageStruct.RegisteredResMessageData{}

	err = json.Unmarshal([]byte(reedPKg_mes.Data), &registeredResMessageData)
	if err != nil {
		fmt.Printf("客户端接收读取服务器发送来的信息反序列化为登录返回消息结构体失败%v\n", err)
		return
	}
	if registeredResMessageData.Code == 200 {
		go ClientKeepsService(coon)
		for {
			ShowLoginMenu(registeredResMessageData.UserName)
		}
	} else {
		fmt.Println(registeredResMessageData.Error)
	}

	return
}

func (this *UserMessageProcessing) Login(userId int, userPwd string) (err error) {
	fmt.Printf("登录输入的账号ID%v密码%v\n", userId, userPwd)
	coon, err := net.Dial("tcp", "0.0.0.0:6666")
	if err != nil {
		fmt.Printf("\"0.0.0.0:6666\"拨号失败%v\n", err)
		return
	}
	defer coon.Close()

	loginMes := messageStruct.LoginMessageData{
		userId,
		userPwd,
		"",
	}

	//开始发送消息

	loginMes_json, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Printf("登录消息序列化失败%v\n", err)
		return
	}

	mes := messageStruct.Message{
		messageStruct.LoginMessageType,
		string(loginMes_json),
	}

	mes_jsonByte, err := json.Marshal(mes)
	if err != nil {
		fmt.Printf("被发送消息结构体序列化为字节切片失败%v\n", err)
		return
	}

	fmt.Printf("客户端发送消息长度%v\n", len(mes_jsonByte))

	mes_TypeByte := [4]byte{}
	binary.BigEndian.PutUint32(mes_TypeByte[0:4], uint32(len(mes_jsonByte)))
	//写入消息的长度
	n, err := coon.Write(mes_TypeByte[0:4])

	if n != 4 || err != nil {
		fmt.Printf("mes_TypeByte消息写入失败%v\n", err)
		return
	}

	//------------------下面是发送的消息内容，上面是内容序列化后的长度

	n, err = coon.Write(mes_jsonByte)
	fmt.Printf("客户端发送消息内容%v\n", string(mes_jsonByte))

	if n != len(mes_jsonByte) || err != nil {
		fmt.Printf("loginMes_Mars消息内容json%v\n", err)
		return
	}

	//-------------------读取客户端返回的消息
	tT := tools.Transfer{
		Conn: coon,
	}
	reedPKg_mes, err := tT.ReedPKg()
	if err != nil {
		if err == io.EOF {
			fmt.Printf("客户端退出%v\n", err)
			return
		} else {
			fmt.Printf("服务端读取失败%v\n", err)
			return
		}
	}
	fmt.Printf("客户端接收读取服务器发送来的信息(结构体)%v\n", reedPKg_mes)
	loginResMes := messageStruct.LoginResMessageData{}

	err = json.Unmarshal([]byte(reedPKg_mes.Data), &loginResMes)
	if err != nil {
		fmt.Printf("客户端接收读取服务器发送来的信息反序列化为登录返回消息结构体失败%v\n", err)
		return
	}
	if loginResMes.Code == 200 {
		go ClientKeepsService(coon)
		for {
			ShowLoginMenu(loginResMes.UserName)
		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}
