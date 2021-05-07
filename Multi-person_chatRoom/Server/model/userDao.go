package model

import (
	"awesomeProject/Multi-person_chatRoom/Common/messageStruct"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var USERDAO *UserDao

func InitUserDao() {
	USERDAO = NewUserDao(pool)
}

type UserDao struct {
	pool *redis.Pool
}

func NewUserDao(Pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: Pool,
	}
	return
}

func (this *UserDao) getUserByid(coon redis.Conn, id int) (user *User, err error) {
	res, err := redis.String(coon.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOT_EXIST
			fmt.Printf("从Redis获取用户失败，用户不存在%v\n", err)
		}
		fmt.Printf("从Redis获取用户失败，未知原因%v\n", err)
		return
	}
	fmt.Printf("根据id从Redis获取用户%v\n", res)
	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Printf("从Redis获取用户信息反序列化为可操作的用户结构体失败%v\n", err)
		return
	}
	fmt.Printf("从Redis获取用户信息反序列化为可操作的用户结构体%v\n", user)
	return
}

func (this UserDao) Registered(registeredMesData *messageStruct.RegisteredMessageData) (err error) {
	coonPool := this.pool.Get()
	defer coonPool.Close()
	_, err = this.getUserByid(coonPool, registeredMesData.UserId)
	if err == redis.ErrNil {
		fmt.Printf("从Redis中用户不存在,可以正常注册%v\n", err)
		err = ERROR_USER_EXIST
		return
	}

	userByte, err := json.Marshal(registeredMesData)
	if err != nil {
		fmt.Printf("注册信息序列化失败%v\n", err)
		return
	}

	_, err = coonPool.Do("HSet", "users", registeredMesData.UserId, string(userByte))
	if err != nil {
		fmt.Printf("注册信息存入数据库失败%v\n", err)
		return
	}
	return
}

func (this *UserDao) LoginVerifi(userId int, userPwd string) (user *User, err error) {
	coonPool := this.pool.Get()
	defer coonPool.Close()
	user, err = this.getUserByid(coonPool, userId)
	if err != nil {
		if err == ERROR_USER_NOT_EXIST {
			ERROR_USER_NOT_EXIST.Error() //用户不存在
			err = ERROR_USER_NOT_EXIST
			return
		}
		fmt.Printf("redis数据库根据id未找到用户%v\n", err)
		return
	}
	if userPwd != user.UserPwd {
		ERROR_USER_PWD.Error() //用户密码不正确
		err = ERROR_USER_PWD
		return
	}
	fmt.Printf("从Redis获取用户信息反序列化为可操作的用户结构体%v\n", user)

	if userPwd != userPwd {
		fmt.Println(ERROR_USER_PWD)
		return
	}
	return
}
