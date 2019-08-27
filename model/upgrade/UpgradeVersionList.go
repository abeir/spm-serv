package upgrade

import (
	"spm-serv/model"
	"spm-serv/model/po"
)

type UpgradeVersionListReq struct {
	model.BaseRequest
	Version string `json:"version"`
	Status string `json:"status"`
}

type UpgradeVersionListRsp struct {
	model.BaseResponse
	Data []*UpgradeVersionListRspData	`json:"data"`
}

type UpgradeVersionListRspData struct {
	po.UpgradeVersion
}
