package spm

import (
	"gopkg.in/go-playground/validator.v8"
	"spm-serv/core"
	"spm-serv/model"
)

//SearchRequest 查询接口请求参数
type SearchRequest struct {
	//包名
	PackageName string	`form:"packageName" binding:"required"`
}

func (s *SearchRequest) GetError(err validator.ValidationErrors) string{
	core.Log.Debugf("SearchRequest GetError: %+v\n", err)
	// 索引对应的是模型的名称和字段
	if val, exist := err["SearchRequest.PackageName"]; exist {
		if val.Field == "PackageName" {
			switch val.Tag{
			case "required":
				return "No package name exists"
			}
		}
	}
	return "Parameter error"
}

//SearchResponse 查询接口返回信息
type SearchResponse struct {
	model.BaseResponse
	//返回数据
	Data []*SearchResponseData		`json:"data"`
}

//SearchResponseData 查询接口数据
type SearchResponseData struct {
	//包名
	Name string 	`json:"name"`
	//描述
	Description string 	`json:"description"`
}