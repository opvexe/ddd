package models

import "time"

// 用户日志实体
type UserLogModel struct {
	*Model
	Id         int
	UserName   string
	LogType    uint8
	LogComment uint8
	Updatetime time.Time
}

func NewUserLogModel(userName string, logType, logComment uint8) *UserLogModel {
	logModel := &UserLogModel{
		UserName:   userName,
		LogType:    logType,
		LogComment: logComment,
	}
	logModel.SetId(logModel.Id)
	logModel.SetName("用户名")
	return logModel
}
