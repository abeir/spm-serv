package pkg

import (
	"spm-serv/model"
	"spm-serv/model/po"
)

type PackageListReq struct {
	model.BaseRequest
	model.Page
	PkgName string 		`json:"pkgName"`
	PkgVersion string 	`json:"pkgVersion"`
}

type PackageListRsp struct {
	model.BaseResponse
	Data []*PackageListRspData `json:"data"`
}

type PackageListRspData struct {
	po.PackageProfile
}
