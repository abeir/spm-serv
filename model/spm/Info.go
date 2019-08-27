package spm

import (
	"gopkg.in/go-playground/validator.v8"
	"spm-serv/core"
	"spm-serv/model"
)

//InfoRequest 详情接口请求参数
type InfoRequest struct {
	//包名
	PackageName string 	`form:"packageName" binding:"required"`
	//版本号
	Version string 		`form:"version"`
}

func (i *InfoRequest) GetError(err validator.ValidationErrors) string{
	core.Log.Debugf("InfoRequest GetError: %+v\n", err)
	// 索引对应的是模型的名称和字段
	if val, exist := err["InfoRequest.PackageName"]; exist {
		if val.Field == "PackageName" {
			switch val.Tag{
			case "required":
				return "No package name exists"
			}
		}
	}
	return "Parameter error"
}

//InfoResponse 详情接口返回信息
type InfoResponse struct {
	model.BaseResponse
	//返回数据
	Data *InfoResponseData	`json:"data"`
}
//InfoResponseData 详情接口返回数据
type InfoResponseData struct {
	Package      Package    `json:"package"`
	Author       Author     `json:"author"`
	Repository   Repository `json:"repository"`
	Version      string     `json:"version"`
	Dependencies []string   `json:"dependencies"`
	PriFilename  string     `json:"priFilename"`
}
