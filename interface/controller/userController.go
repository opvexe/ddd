package controller

import (
	"github.com/gin-gonic/gin"
	"github.xiadat.com/goddd/application/dto"
	"github.xiadat.com/goddd/application/services"
)

type UserController struct {
	UserSvc *services.UserService
}

func (this *UserController) UserDetail(ctx *gin.Context) {
	simpleUserReq := &dto.SimpleUserReq{}
	ctx.ShouldBind(simpleUserReq)
	this.UserSvc.GetSimpleUserInfo(simpleUserReq)
}
