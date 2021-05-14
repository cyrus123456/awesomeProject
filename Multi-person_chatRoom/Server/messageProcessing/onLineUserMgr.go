package messageProcessing

import "fmt"

type OnLineUserMgr struct {
	onLineUsers map[int]*UserMessageProcessing
}

//服务端唯一并且很多地方会用到
var onLineUserMgr *OnLineUserMgr

func InitOnLineUserMgr() {
	onLineUserMgr = &OnLineUserMgr{
		onLineUsers: make(map[int]*UserMessageProcessing, 1024),
	}
}

//----------------------------------赠删改查
//增加以及修改
func (this *OnLineUserMgr) AddOnlineUser(userMessageProcessing *UserMessageProcessing) {
	this.onLineUsers[userMessageProcessing.UserId] = userMessageProcessing
}

//删除
func (this *OnLineUserMgr) DelOnlineUser(userId int) {
	delete(this.onLineUsers, userId)
}

//-------------------------------查询
//全部在线用户
func (this *OnLineUserMgr) AllOnlineUser() map[int]*UserMessageProcessing {
	return this.onLineUsers
}

//根据ID查找
func (this *OnLineUserMgr) FindOnlineUserById(userId int) (userMessageProcessing *UserMessageProcessing, err error) {
	userMessageProcessing, ok := this.onLineUsers[userId]
	if !ok {
		err = fmt.Errorf("当前用户不在线%v\n", userId)
	}
	return
}
