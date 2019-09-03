package console

import (
	"gopkg.in/go-playground/validator.v8"
	"spm-serv/core"
	"spm-serv/model"
)

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (l *LoginReq) GetError(err validator.ValidationErrors) string{
	core.Log.Debugf("LoginReq GetError: %+v\n", err)
	// 索引对应的是模型的名称和字段
	if val, exist := err["LoginReq.Username"]; exist {
		if val.Field == "Username" {
			switch val.Tag{
			case "required":
				return "Username is empty"
			}
		}
		if val.Field == "Password" {
			switch val.Tag{
			case "required":
				return "Password is empty"
			}
		}
	}
	return "Parameter error"
}

type LoginRsp struct {
	model.BaseResponse
	Data *LoginRspData  `json:"data"`
}

type LoginRspData struct {
	Token string 		`json:"token"`
}


