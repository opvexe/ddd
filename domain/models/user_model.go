package models

import (
	"crypto/md5"
	"fmt"

	"github.xiadat.com/goddd/domain/valueobjs"
)

// 定义用户实体
// 有唯一标识
// 包含属性
// 数据校验，操作前置函数
type UserModel struct {
	*Model   //实体名称
	UserId   int
	UserName string
	UserPwd  string

	UserExtra *valueobjs.UserExtra // 2.值对象
}

func NewUserModel(attrs ...UserAttrFunc) *UserModel {
	user := &UserModel{}
	UserAttrFuncs(attrs).apply(user)
	user.Model = &Model{}
	user.SetName("User Entiry")
	user.SetId(user.UserId)
	return user
}

func (this *UserModel) BeforeSave() {
	// md5
	this.UserPwd = fmt.Sprintf("%x", md5.Sum([]byte(this.UserPwd)))
}
