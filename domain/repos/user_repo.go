package repos

import "github.xiadat.com/goddd/domain/models"

type IUserRepo interface {
	FindByName(name string) *models.UserModel
	SveUser(*models.UserModel) error
	UpdateUser(*models.UserModel) error
	DeleteUser(*models.UserModel) error
}

type IUserLoginRepo interface {
	FindByName(name string) *models.UserLogModel
	SveUser(*models.UserLogModel) error
}
