package pkg

import (
	"gopkg.in/go-playground/validator.v8"
	"spm-serv/core"
	"spm-serv/model"
	"spm-serv/model/po"
)

type PackageInfoReq struct {
	model.BaseRequest
	Id string `json:"id" binding:"required"`
}

func (p *PackageInfoReq) GetError(err validator.ValidationErrors) string{
	core.Log.Debugf("PackageInfoReq GetError: %+v\n", err)
	// 索引对应的是模型的名称和字段
	if val, exist := err["PackageInfoReq.Id"]; exist {
		if val.Field == "Id" {
			switch val.Tag{
			case "required":
				return "ID is empty"
			}
		}
	}
	return "Parameter error"
}

type PackageInfoRsp struct {
	model.BaseResponse
	Data *PackageInfoRspData
}

type PackageInfoRspData struct {
	po.PackageProfile
}
