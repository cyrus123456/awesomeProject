package messageProcessing

import (
	"awesomeProject/Multi-person_chatRoom/Common/messageStruct"
	"awesomeProject/Multi-person_chatRoom/Server/model"
	"awesomeProject/Multi-person_chatRoom/Server/tools"
	"encoding/json"
	"fmt"
	"net"
)

type UserMessageProcessing struct {
	Coon net.Conn
}

//处理注册消息
func (this UserMessageProcessing) Registered_messageProcessing(mes *messageStruct.Message) (err error) {
	//实例化一个空的注册消息结构体
	registeredMesData := messageStruct.RegisteredMessageData{}

	//把服务端接收到的序列化的登录消息反序列化然后存到&registeredMesData结构体
	err = json.Unmarshal([]byte(mes.Data), &registeredMesData)
	if err != nil {
		fmt.Printf("服务器处理注册信息反序列化存入登录结构体失败%v\n", err)
		return
	}
	fmt.Printf("服务端验收到的登录消息内容%v\n", registeredMesData)

	//判断账号id是否重复
	//实例化一个空的注册消息返回结构体
	registeredResMessageData := messageStruct.RegisteredResMessageData{}
	//注册验证
	err = model.USERDAO.Registered(&registeredMesData)
	if err != nil {
		if err == model.ERROR_USER_EXIST {
			registeredResMessageData.Code = 500 //用户已存在
			registeredResMessageData.Error = model.ERROR_USER_EXIST.Error()
			fmt.Printf("用户已存在%v\n", err)
		} else {
			registeredResMessageData.Code = 505 //未知错误
			registeredResMessageData.Error = "未知错误"
			fmt.Printf("用户注册失败,未知错误%v\n", err)
			return
		}
	} else {
		registeredResMessageData.Code = 200 //用户用户注册成功
		registeredResMessageData.UserName = registeredMesData.UserName
		fmt.Printf("用户已存在%v\n", err)
	}
	//-------------------重新序列化网络传输返回客户端
	loginResMes_jsonByte, err := json.Marshal(registeredResMessageData)
	if err != nil {
		fmt.Printf("返回注册消息结构体序列化失败%v\n", err)
		return
	}
	fmt.Printf("返回注册消息结构体序列化%v\n", registeredResMessageData)

	mes.Data = string(loginResMes_jsonByte) //序列化以后才能转成字符串赋值给
	mes.Type = messageStruct.RegisteredResMesRegisteredsageType
	//将消息结构体序列化成字节切片准备发送
	mes_jsonByte, err := json.Marshal(mes)
	if err != nil {
		fmt.Printf("将要返回给客户端消息结构体序列化为字节切片失败%v\n", err)
		return
	}
	fmt.Printf("返回消息结构体序列化成功%v\n", mes_jsonByte)

	//向返回客户端消息

	tT := &tools.Transfer{
		Conn: this.Coon,
	}

	mesStr, err := tT.WritePKg(mes_jsonByte)
	if err != nil {
		fmt.Printf("将要返回给客户端消息发送失败%v\n", err)
		return
	}
	fmt.Printf("将要返回给客户端消息发送成功%v\n", mesStr)

	return
}

//处理登陆消息
func (this *UserMessageProcessing) Login_messageProcessing(mes *messageStruct.Message) (err error) {
	//实例化一个空的登录消息结构体
	loginMes := messageStruct.LoginMessageData{}

	//把服务端接收到的序列化的登录消息反序列化然后存到&loginMes结构体
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Printf("处理登录信息反序列化存入登录结构体失败%v\n", err)
		return
	}
	fmt.Printf("服务端验收到的登录消息内容%v\n", loginMes)

	//判断密码账号是否正确
	//mes.Type = messageStruct.LoginResMessageType
	loginResMes := messageStruct.LoginResMessageData{}

	user, err := model.USERDAO.LoginVerifi(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOT_EXIST {
			loginResMes.Code = 500 //账号密码错误
			loginResMes.Error = model.ERROR_USER_NOT_EXIST.Error()
			fmt.Printf("服务端验证客户失败、未注册%v\n", err)
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 400 //账号密码错误
			loginResMes.Error = model.ERROR_USER_PWD.Error()
			fmt.Printf("服务端验证客户失败、密码错误%v\n", err)
		} else {
			loginResMes.Code = 505 //账号密码错误
			loginResMes.Error = err.Error()
			fmt.Printf("服务端验证客户失败、未知错误%v\n", err)
		}
	} else {
		loginResMes.Code = 200 //200表示通过
		loginResMes.UserName = user.UserName
		fmt.Printf("服务端验证用户%v登陆成功\n", user.UserName)
	}

	loginResMes_jsonByte, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Printf("登录返回消息结构体序列化失败%v\n", err)
		return
	}
	fmt.Printf("登录返回消息结构体序列化%v\n", loginResMes_jsonByte)

	mes.Data = string(loginResMes_jsonByte) //序列化以后才能转成字符串赋值给
	mes.Type = messageStruct.LoginResMessageType
	//将消息结构体序列化成字节切片准备发送
	mes_jsonByte, err := json.Marshal(mes)
	if err != nil {
		fmt.Printf("将要返回给客户端消息结构体序列化为字节切片失败%v\n", err)
		return
	}
	fmt.Printf("返回消息结构体序列化成功%v\n", mes_jsonByte)

	//向返回客户端消息

	tT := &tools.Transfer{
		Conn: this.Coon,
	}

	mesStr, err := tT.WritePKg(mes_jsonByte)
	if err != nil {
		fmt.Printf("将要返回给客户端消息发送失败%v\n", err)
		return
	}
	fmt.Printf("将要返回给客户端消息发送成功%v\n", mesStr)

	return
}
