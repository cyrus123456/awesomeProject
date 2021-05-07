package tools

import (
	"awesomeProject/Multi-person_chatRoom/Common/messageStruct"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
}

//解析读取消息
func (this *Transfer) ReedPKg() (mes messageStruct.Message, err error) {
	var Mes_TypeByte [8096]byte

	_, err = this.Conn.Read(Mes_TypeByte[0:4])

	if err != nil {
		fmt.Printf("服务器读取客户端发送来的信息长度mes_TypeByte失败%v\n", err)
		return
	}
	var Mes_TypeByte_Len uint32

	Mes_TypeByte_Len = binary.BigEndian.Uint32(Mes_TypeByte[0:4])
	fmt.Printf("服务器读取客户端发送来的信息长度%v\n", Mes_TypeByte_Len)

	//----------------以下信息处理发来的消息内容，上面是消息长度

	fmt.Printf("服务器读取客户端发送来的信息（没有反序列化）是%v\n", Mes_TypeByte[0:Mes_TypeByte_Len])

	n, err := this.Conn.Read(Mes_TypeByte[0:Mes_TypeByte_Len])

	//if err != nil {
	//	fmt.Printf("mes_TypeByte读取失败%v\n", err)
	//	return
	//}

	if n != int(Mes_TypeByte_Len) || err != nil {
		fmt.Printf("服务器对客户端发送来的信息读取失败%v\n", err)
		return
	}

	err = json.Unmarshal(Mes_TypeByte[:Mes_TypeByte_Len], &mes)
	if err != nil {
		fmt.Printf("mes_TypeByte读取失败%v\n", err)
		return
	}

	return
}

//发送消息
func (this *Transfer) WritePKg(mes_jsonByte []byte) (mes messageStruct.Message, err error) {
	var Mes_TypeByte [8096]byte

	binary.BigEndian.PutUint32(Mes_TypeByte[0:4], uint32(len(mes_jsonByte)))
	n, err := this.Conn.Write(Mes_TypeByte[:4])
	if n != 4 || err != nil {
		fmt.Printf("服务器返回登录消息长度写入发送失败%v\n", err)
		return
	}
	fmt.Printf("服务器返回登录消息长度%v\n", n)

	n, err = this.Conn.Write(mes_jsonByte)
	if n != len(mes_jsonByte) || err != nil {
		fmt.Printf("服务器返回登录消息内容失败%v\n", err)
		return
	}
	fmt.Printf("服务器返回登录消息内容成功%v\n", len(mes_jsonByte))

	return
}
