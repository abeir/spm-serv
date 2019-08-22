package model

import (
	"gopkg.in/go-playground/validator.v8"
	"spm-serv/core"
)

//Package 包信息
type Package struct {
	//包名
	Name string `json:"name" binding:"required"`
	//描述
	Description string `json:"description"`
}
//Author 作者信息
type Author struct {
	//姓名
	Name string 	`json:"name"`
	//邮箱
	Email string	`json:"email"`
	//描述
	Description string `json:"description"`
}
//Repository 仓库信息
type Repository struct {
	//仓库url
	Url string 		`json:"url" binding:"required"`
}

//PublishRequest 推送接口请求参数
type PublishRequest struct {
	Package Package		`json:"package"`
	Author Author		`json:"author"`
	Repository Repository	`json:"repository"`
	Version string		`json:"version" binding:"required,checkVersion"`
	Dependencies []string	`json:"dependencies"`
	PriFilename	string 	`json:"priFilename" binding:"required"`
	Force	string		`json:"force"`
}

func (p *PublishRequest) GetError(err validator.ValidationErrors) string{
	core.Log.Debugf("Publish GetError: %+v\n", err)


	// 索引对应的是模型的名称和字段
	if val, exist := err["PublishRequest.Package.Name"]; exist {
		if val.Field == "Name" {
			switch val.Tag{
			case "required":
				return "No package name exists"
			}
		}
	}
	if val, exist := err["PublishRequest.Repository.Url"]; exist {
		if val.Field == "Url" {
			switch val.Tag{
			case "required":
				return "Repository URL does not exist"
			}
		}
	}
	if val, exist := err["PublishRequest.Version"]; exist {
		if val.Field == "Version" {
			switch val.Tag{
			case "required":
				return "Version number does not exist"
			case "checkVersion":
				return "Version format error, e.g. 1.0.0"
			}
		}
	}
	if val, exist := err["PublishRequest.PriFilename"]; exist {
		if val.Field == "PriFilename" {
			switch val.Tag{
			case "required":
				return "Pri file name does not exist"
			}
		}
	}
	return "Parameter error"
}


//PublishResponse 推送接口返回信息
type PublishResponse struct {
	BaseResponse
}
