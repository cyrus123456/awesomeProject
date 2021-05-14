package main

import (
	"awesomeProject/Multi-person_chatRoom/Server/messageProcessing"
	"awesomeProject/Multi-person_chatRoom/Server/model"
	"fmt"
	"net"
	"time"
)

////解析读取消息
//func reedPKg(conn net.Conn) (mes messageStruct.Message, err error) {
//
//	mes_TypeByte := make([]byte, 8096)
//	_, err = conn.Read(mes_TypeByte[0:4])
//
//	if err != nil {
//		fmt.Printf("服务器读取客户端发送来的信息长度mes_TypeByte失败%v\n", err)
//		return
//	}
//
//	var mes_TypeByte_Len uint32
//	mes_TypeByte_Len = binary.BigEndian.Uint32(mes_TypeByte[0:4])
//	fmt.Printf("服务器读取客户端发送来的信息长度%v\n", mes_TypeByte_Len)
//
//	//----------------以下信息处理发来的消息内容，上面是消息长度
//
//	fmt.Printf("服务器读取客户端发送来的信息（没有反序列化）是%v\n", mes_TypeByte[0:mes_TypeByte_Len])
//
//	n, err := conn.Read(mes_TypeByte[0:mes_TypeByte_Len])
//
//	//if err != nil {
//	//	fmt.Printf("mes_TypeByte读取失败%v\n", err)
//	//	return
//	//}
//
//	if n != int(mes_TypeByte_Len) || err != nil {
//		fmt.Printf("服务器对客户端发送来的信息读取失败%v\n", err)
//		return
//	}
//
//	err = json.Unmarshal(mes_TypeByte[:mes_TypeByte_Len], &mes)
//	if err != nil {
//		fmt.Printf("mes_TypeByte读取失败%v\n", err)
//		return
//	}
//
//	return
//}
//
////发送消息
//func writePKg(conn net.Conn, mes_jsonByte []byte) (mes messageStruct.Message, err error) {
//	mes_TypeByte := [4]byte{}
//	binary.BigEndian.PutUint32(mes_TypeByte[0:4], uint32(len(mes_jsonByte)))
//	n, err := conn.Write(mes_TypeByte[:4])
//	if n != 4 || err != nil {
//		fmt.Printf("服务器返回登录消息长度写入发送失败%v\n", err)
//		return
//	}
//	fmt.Printf("服务器返回登录消息长度%v\n", n)
//
//
//	n, err = conn.Write(mes_jsonByte)
//	if n != len(mes_jsonByte) || err != nil {
//		fmt.Printf("服务器返回登录消息内容失败%v\n", err)
//		return
//	}
//	fmt.Printf("服务器返回登录消息内容成功%v\n", len(mes_jsonByte))
//
//	return
//}

//处理登陆消息
//func Login_messageProcessing(coon net.Conn, mes *messageStruct.Message) (err error) {
//	loginMes := messageStruct.LoginMessageData{}
//
//	//把mes.Data序列化然后存到&loginMes结构体
//	err = json.Unmarshal([]byte(mes.Data), &loginMes)
//
//	if err != nil {
//		fmt.Printf("处理登录信息反序列化存入登录结构体失败%v\n", err)
//		return
//	}
//
//	//判断密码账号是否正确
//	mes.Type = messageStruct.LoginResMessageType
//	loginResMes := messageStruct.LoginResMessageData{}
//	if loginMes.UserId == 100 && loginMes.UserPwd == 123456 {
//		loginResMes.Code = 200 //200表示通过
//	} else {
//		loginResMes.Code = 500 //账号密码错误
//		loginResMes.Error = "账号或者密码错误！"
//	}
//
//	loginResMes_jsonByte, err := json.Marshal(loginResMes)
//	if err != nil {
//		fmt.Printf("登录返回消息结构体序列化失败%v\n", err)
//		return
//	}
//	fmt.Printf("登录返回消息结构体序列化%v\n", loginResMes_jsonByte)
//
//	mes.Data = string(loginResMes_jsonByte) //序列化以后才能转成字符串赋值给
//	//将消息结构体序列化成字节切片准备发送
//	mes_jsonByte, err := json.Marshal(mes)
//	if err != nil {
//		fmt.Printf("将要返回给客户端消息结构体序列化为字节切片失败%v\n", err)
//		return
//	}
//	fmt.Printf("返回消息结构体序列化成功%v\n", mes_jsonByte)
//
//	//向返回客户端消息
//
//	tools.WritePKg(coon, mes_jsonByte)
//
//	return
//}

//根据不同消息类型处理不同消息
//func ClassifiedProcessingMessages(coon net.Conn, mes *messageStruct.Message) (err error) {
//	switch mes.Type {
//
//	case messageStruct.LoginMessageType:
//		fmt.Println("处理登陆消息")
//		err = messageProcessing.Login_messageProcessing(coon, mes)
//		if err != nil {
//			fmt.Printf("处理登陆消息失败%v\n", err)
//			return
//		}
//	default:
//		fmt.Println("消息类型错误")
//
//	}
//	return
//}

//处理客户端消息
func communication(conn net.Conn) {
	defer conn.Close()

	ClassifiedProcessing_Messages := &messageProcessing.ClassifiedProcessing_Messages{
		Coon: conn,
	}
	//for {
	err := ClassifiedProcessing_Messages.CommunicationFor()
	if err != nil {
		fmt.Printf("循环读取客户端发送消息失败%v\n", err)
		return
	}
	//}

	//for{
	//mes, err := tools.Transfer.ReedPKg(conn)
	//if err != nil {
	//	if err == io.EOF {
	//		fmt.Printf("客户端退出%v\n", err)
	//		return
	//	} else {
	//		fmt.Printf("服务端读取失败%v\n", err)
	//		return
	//	}
	//}
	//fmt.Printf("服务器对客户端发送来的信息(反序列化以后)%v\n", mes)
	//ClassifiedProcessing_messages := ClassifiedProcessing_Messages{
	//	Coon: conn,
	//}
	//err = ClassifiedProcessing_messages.ClassifiedProcessingMessages(&mes)
	//if err != nil {
	//	fmt.Printf("服务器信息结构体分类处理失败%v\n", err)
	//	return
	//
	//}
	//}

}

func main() {
	//初始化连接池
	model.InitPool(16, 0, time.Second, "127.0.0.1:6379")
	model.InitUserDao()
	//初始化在线用户
	messageProcessing.InitOnLineUserMgr()
	//客户端提示信息
	fmt.Println("服务器监听端口6666……")
	listener, err := net.Listen("tcp", "0.0.0.0:6666")
	defer listener.Close()
	if err != nil {
		fmt.Printf("\"0.0.0.0:6666\"端口监听失败%v\n", err)
		return
	} else {
		fmt.Println("\"0.0.0.0:6666\"端口监听成功")
	}
	//监听成功持续链接
	for {
		fmt.Println("等待客户端消息")

		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("接收消息失败%v\n", err)
		}

		go communication(conn)
	}

}
