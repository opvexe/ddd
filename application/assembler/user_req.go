package assembler

import (
	"github.com/go-playground/validator/v10"
	"github.xiadat.com/goddd/application/dto"
	"github.xiadat.com/goddd/domain/models"
)

type UserReq struct {
	v *validator.Validate
}

func (this *UserReq) D2M_UserModel(dto *dto.SimpleUserReq) *models.UserModel {
	if err := this.v.Struct(dto); err != nil {
		panic(err)
	}
	return models.NewUserModel(
		models.WithUserId(dto.Id),
	)

}
