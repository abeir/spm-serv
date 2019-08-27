package service

import (
	"spm-serv/model/upgrade"
)

type UpgradeVersionService struct {
}

//查询列表
func (u *UpgradeVersionService) List(req *upgrade.UpgradeVersionListReq) (*upgrade.UpgradeVersionListRsp, error){
	return nil, nil
}

//查询详情
func (u *UpgradeVersionService) Info(req *upgrade.UpgradeVersionInfoReq) (*upgrade.UpgradeVersionInfoRsp, error){
	return nil, nil
}

//上传spm
func (u *UpgradeVersionService) Upload(req *upgrade.UpgradeVersionUploadReq) (*upgrade.UpgradeVersionUploadRsp, error){
	return nil, nil
}

//发布spm
func (u *UpgradeVersionService) Release(req *upgrade.UpgradeVersionStatusReq) (*upgrade.UpgradeVersionStatusRsp, error){
	return nil, nil
}

//下架spm
func (u *UpgradeVersionService) Detain(req *upgrade.UpgradeVersionStatusReq) (*upgrade.UpgradeVersionStatusRsp, error){
	return nil, nil
}
