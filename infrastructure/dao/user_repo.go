package dao

// // infrastructure 基础实施曾
// application  :应用曾 主要就是连接domain 和interface曾
// request -> 应用曾 -> 领域曾 -> 基础实施曾
import (
	"github.com/jinzhu/gorm"
	"github.xiadat.com/goddd/domain/models"
	"github.xiadat.com/goddd/domain/repos"
)

var _ repos.IUserRepo = &UserRepo{}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (this *UserRepo) FindByName(name string) *models.UserModel {
	user := &models.UserModel{}
	if err := this.db.Where("user_name=?", name).Find(user).Error; err != nil {
		return nil
	}
	return user
}

func (this *UserRepo) SveUser(*models.UserModel) error {
	return nil
}

func (this *UserRepo) UpdateUser(*models.UserModel) error {
	return nil
}

func (this *UserRepo) DeleteUser(*models.UserModel) error {
	return nil
}
