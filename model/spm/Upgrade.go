package spm

import (
	"gopkg.in/go-playground/validator.v8"
	"spm-serv/core"
	"spm-serv/model"
)

type UpgradeRequest struct {
	//版本号
	Version string	`form:"version" binding:"required,checkVersion"`
}

func (p *UpgradeRequest) GetError(err validator.ValidationErrors) string{
	core.Log.Debugf("UpgradeRequest GetError: %+v\n", err)
	if val, exist := err["UpgradeRequest.Version"]; exist {
		if val.Field == "Version" {
			switch val.Tag{
			case "required":
				return "Version number does not exist"
			case "checkVersion":
				return "Version format error, e.g. 1.0.0"
			}
		}
	}
	return "Parameter error"
}

//UpgradeResponse 查询新版本
type UpgradeResponse struct {
	model.BaseResponse
	//返回数据
	Data *UpgradeResponseData	`json:"data"`
}
//UpgradeResponseData 查询版本返回数据
type UpgradeResponseData struct {
	Version string 	`json:"version"`
}
