package service

import (
	"spm-serv/core"
	"spm-serv/dao"
	"spm-serv/model"
)

//查询新版本
func UpgradeService(req *model.UpgradeRequest) (*model.UpgradeResponse, error){
	core.Log.Debugf("UpgradeService request: %+v\n", req)

	upgradeVersion, err := dao.UpgradeVersionDaoImpl.SelectLatestVersion()
	if err!=nil {
		return nil, err
	}

	core.Log.Debugf("upgradeVersion %+v\n", upgradeVersion)

	rsp := &model.UpgradeResponse{}
	//数据库中没有记录
	if upgradeVersion.IsEmpty() {
		rsp.Code = model.CODE_ERROR
		rsp.Msg = "It's the latest version"
		return rsp, nil
	}
	//判断数据库中版本与上传的版本号大小
	verInDb, err := core.NewVersion(upgradeVersion.Version)
	if err!=nil {
		return nil, err
	}
	verInForm, err := core.NewVersion(req.Version)
	if err!=nil {
		return nil, err
	}
	if verInDb.Compare(verInForm) > 0 {
		rsp.Code = model.CODE_SUCCESS
		rsp.Msg = "There is a new version"
		rsp.Data = &model.UpgradeResponseData{Version:upgradeVersion.Version}
		return rsp, nil
	}

	rsp.Code = model.CODE_ERROR
	rsp.Msg = "It's the latest version"
	return rsp, nil
}

//下载版本，返回待下载的文件路径
func DownloadVersion(req *model.UpgradeRequest) (string, error){
	core.Log.Debugf("DownloadVersion request: %+v\n", req)

	upgradeVersion, err := dao.UpgradeVersionDaoImpl.SelectByVersion(req.Version)
	if err!=nil {
		return "", err
	}
	core.Log.Debugf("upgradeVersion: %+v\n", upgradeVersion)
	return upgradeVersion.Path, nil
}