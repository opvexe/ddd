package services

import (
	"github.xiadat.com/goddd/application/assembler"
	"github.xiadat.com/goddd/application/dto"
	"github.xiadat.com/goddd/domain/repos"
)

type UserService struct {
	assUerReq *assembler.UserReq
	assUsrRsp *assembler.UserRsp
	userRepo  repos.IUserRepo
}

func (this *UserService) GetSimpleUserInfo(req *dto.SimpleUserReq) *dto.SimpleUserInfo {
	userModel := this.assUerReq.D2M_UserModel(req)
	this.userRepo.FindByName(userModel.Name)
	return this.assUsrRsp.M2d_SimpleUserINfo(userModel)
}
