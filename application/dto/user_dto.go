package dto

import "time"

////////////////////////////////////////////////////////
// 用户请求对象
//

type SimpleUserReq struct {
	Id int `uri:"id" binding:"required,min=100"`
}

///////////////////////////////////////////////////////////
// 输出对象
/// 以下都是

type SimpleUserInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

type UserInfo struct {
	Id    int        `json:"id"`
	Name  string     `json:"name"`
	City  string     `json:"city"`
	Phone string     `json:"phone"`
	Logs  []*UserLog `json:"logs"`
}

type UserLog struct {
	Id   int       `json:"id"`
	Log  string    `json:"log"`
	Data time.Time `json:"date"`
}
