package upgrade

import (
	"gopkg.in/go-playground/validator.v8"
	"spm-serv/core"
	"spm-serv/model"
	"spm-serv/model/po"
)

type UpgradeVersionInfoReq struct {
	model.BaseRequest
	Id string `json:"id" binding:"required"`
}

func (u *UpgradeVersionInfoReq) GetError(err validator.ValidationErrors) string{
	core.Log.Debugf("UpgradeVersionInfoReq GetError: %+v\n", err)
	// 索引对应的是模型的名称和字段
	if val, exist := err["UpgradeVersionInfoReq.Id"]; exist {
		if val.Field == "Id" {
			switch val.Tag{
			case "required":
				return "ID is empty"
			}
		}
	}
	return "Parameter error"
}

type UpgradeVersionInfoRsp struct {
	model.BaseResponse
	Data *UpgradeVersionInfoData `json:"data"`
}

type UpgradeVersionInfoData struct {
	po.UpgradeVersion
}