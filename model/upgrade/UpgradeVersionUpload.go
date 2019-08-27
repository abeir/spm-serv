package upgrade

import (
	"gopkg.in/go-playground/validator.v8"
	"spm-serv/core"
	"spm-serv/model"
	"spm-serv/model/po"
)

type UpgradeVersionUploadReq struct {
	model.BaseRequest

	Version string 		`json:"version" binding:"required,checkVersion"`
	Description string 	`json:"description"`
}

func (u *UpgradeVersionUploadReq) GetError(err validator.ValidationErrors) string{
	core.Log.Debugf("UpgradeVersionUploadReq GetError: %+v\n", err)
	// 索引对应的是模型的名称和字段
	if val, exist := err["UpgradeVersionUploadReq.Version"]; exist {
		if val.Field == "Version" {
			switch val.Tag{
			case "required":
				return "ID is empty"
			case "checkVersion":
				return "Version format error, e.g. 1.0.0"
			}
		}
	}
	return "Parameter error"
}

type UpgradeVersionUploadRsp struct {
	model.BaseResponse
	Data *UpgradeVersionUploadRspData `json:"data"`
}

type UpgradeVersionUploadRspData struct {
	po.UpgradeVersion
}
