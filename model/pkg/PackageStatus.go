package pkg

import (
	"gopkg.in/go-playground/validator.v8"
	"spm-serv/core"
	"spm-serv/model"
	"spm-serv/model/po"
)

type PackageStatusReq struct {
	model.BaseRequest
	Id string `json:"id" binding:"required"`
	Status string `json:"status" binding:"required,checkValues=[0,1]"`
}

func (p *PackageStatusReq) GetError(err validator.ValidationErrors) string{
	core.Log.Debugf("PackageStatusReq GetError: %+v\n", err)
	// 索引对应的是模型的名称和字段
	if val, exist := err["PackageStatusReq.Id"]; exist {
		if val.Field == "Id" {
			switch val.Tag{
			case "required":
				return "ID is empty"
			}
		}
	}
	if val, exist := err["PackageStatusReq.Status"]; exist {
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

type PackageStatusRsp struct {
	model.BaseResponse
	Data *PackageStatusRspData `json:"data"`
}

type PackageStatusRspData struct {
	po.PackageProfile
}
