package service

import (
	"errors"

	"github.xiadat.com/goddd/domain/repos"
)

// service  向上 interface 曾调用
// 向下调用领域曾

type UserLoginService struct {
	userRepo repos.IUserRepo
}

func (this *UserLoginService) UserLogin(username string, userPwd string) (string, error) {
	user := this.userRepo.FindByName(username)
	if user.UserId > 0 {
		// md5
		if user.UserPwd == userPwd {
			return "1000200", nil
		}
		return "1000404", errors.New("password wrong")
	} else {
		return "1000404", errors.New("user not exists")
	}
}
