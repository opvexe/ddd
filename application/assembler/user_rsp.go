package assembler

// assembler装配
// 主要完成的是领域对象和相应的dto对象之间的类型转换
// 数据填充，多个领域对象的组装成一个dto

import (
	"github.xiadat.com/goddd/application/dto"
	"github.xiadat.com/goddd/domain/aggregates"
	"github.xiadat.com/goddd/domain/models"
)

type UserRsp struct {
}

// 把用户实体映射为简单用户的Dto
func (this *UserRsp) M2d_SimpleUserINfo(user *models.UserModel) *dto.SimpleUserInfo {
	simpleUser := &dto.SimpleUserInfo{}
	simpleUser.Id = user.Id
	simpleUser.Name = user.UserName
	simpleUser.City = user.UserExtra.UserCity
	return simpleUser
}

func (this *UserRsp) M2D_UserInfo(mem *aggregates.Member) *dto.UserInfo {
	userInfo := &dto.UserInfo{}
	userInfo.Id = mem.Log.Id
	userInfo.Logs = this.M2D_UserLogs(mem.GetLogs())
	return userInfo
}

func (this *UserRsp) M2D_UserLogs(logs []*models.UserLogModel) []*dto.UserLog {
	return nil
}
