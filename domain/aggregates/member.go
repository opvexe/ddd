package aggregates

import (
	"github.xiadat.com/goddd/domain/models"
	"github.xiadat.com/goddd/domain/repos"
)

// 会员聚会
// user + log
type Member struct {
	User        *models.UserModel
	Log         *models.UserLogModel
	userRepo    repos.IUserRepo
	userLogRepo repos.IUserLoginRepo
}

func NewMember(user *models.UserModel, userRepo repos.IUserRepo, userLogRepo repos.IUserLoginRepo) *Member {
	return &Member{
		User:        user,
		userRepo:    userRepo,
		userLogRepo: userLogRepo,
	}
}

func NewMemberById(id string, userRepo repos.IUserRepo, userLogRepo repos.IUserLoginRepo) *Member {
	user := userRepo.FindByName(id)
	return &Member{
		User:        user,
		userRepo:    userRepo,
		userLogRepo: userLogRepo,
	}
}

// 创建会员
func (this *Member) Create() error {
	if err := this.userRepo.SveUser(this.User); err != nil {
		return err
	}
	this.Log = models.NewUserLogModel(
		this.User.UserName,
		1,
		1,
	)
	return this.userLogRepo.SveUser(this.Log)
}

func (this *Member) GetLogs() []*models.UserLogModel {

	return nil
}
