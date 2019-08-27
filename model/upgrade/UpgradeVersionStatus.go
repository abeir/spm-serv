package upgrade

import (
	"gopkg.in/go-playground/validator.v8"
	"spm-serv/core"
	"spm-serv/model"
	"spm-serv/model/po"
)

type UpgradeVersionStatusReq struct {
	model.BaseRequest
	Id string `json:"id" binding:"required"`
	Status string `json:"status" binding:"required,checkValues=[0,1,2]"`
}

func (u *UpgradeVersionStatusReq) GetError(err validator.ValidationErrors) string{
	core.Log.Debugf("UpgradeVersionStatusReq GetError: %+v\n", err)
	// 索引对应的是模型的名称和字段
	if val, exist := err["UpgradeVersionStatusReq.Id"]; exist {
		if val.Field == "Id" {
			switch val.Tag{
			case "required":
				return "ID is empty"
			}
		}
	}
	if val, exist := err["UpgradeVersionStatusReq.Status"]; exist {
		if val.Field == "Status" {
			switch val.Tag{
			case "required":
				return "Status is empty"
			case "checkValues":
				return "Invalid status value"
			}
		}
	}
	return "Parameter error"
}

type UpgradeVersionStatusRsp struct {
	model.BaseResponse
	Data *UpgradeVersionStatusRspData `json:"data"`
}

type UpgradeVersionStatusRspData struct {
	po.UpgradeVersion
}
