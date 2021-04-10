package models

import "github.xiadat.com/goddd/domain/valueobjs"

type UserAttrFunc func(model *UserModel) // 设置User的属性
type UserAttrFuncs []UserAttrFunc

//apply 方法
func (this UserAttrFuncs) apply(u *UserModel) {
	for _, f := range this {
		f(u)
	}
}

// 传参数
func WithUserId(id int) UserAttrFunc {
	return func(model *UserModel) {
		model.UserId = id
	}
}

func WithUserName(userName string) UserAttrFunc {
	return func(model *UserModel) {
		model.UserName = userName
	}
}

func WithUserPassword(userPwd string) UserAttrFunc {
	return func(model *UserModel) {
		model.UserPwd = userPwd
	}
}

func WithUserExtra(extra *valueobjs.UserExtra) UserAttrFunc {
	return func(model *UserModel) {
		model.UserExtra = extra
	}
}
