package service

import (
	"spm-serv/core"
	"spm-serv/dao"
	"spm-serv/model"
)

func InfoService(req *model.InfoRequest) (*model.InfoResponse, error){
	core.Log.Debugf("InfoService request: %+v\n", *req)

	pkgVersion := req.Version
	if req.Version=="" {
		lastVersion, err := dao.LastVersionDaoImpl.SelectByPkgName(req.PackageName)
		if err!=nil {
			return nil, err
		}
		pkgVersion = lastVersion.PkgVersion
	}

	pkgProfile, err := dao.PackageProfileDaoImpl.SelectByPkgNameAndPkgVersion(req.PackageName, pkgVersion)
	if err!=nil {
		return nil, err
	}
	rsp := &model.InfoResponse{
		BaseResponse: model.BaseResponse{
			Code: model.CODE_SUCCESS,
			Msg: "ok",
		},
	}
	if pkgProfile.IsEmpty() {
		return rsp, nil
	}

	rsp.Data = &model.InfoResponseData{
			Package:      model.Package{
				Name: pkgProfile.PkgName,
				Description: pkgProfile.PkgDesc,
			},
			Author:       model.Author{
				Name: pkgProfile.AuthorName,
				Email: pkgProfile.AuthorEmail,
				Description: pkgProfile.AuthorDesc,
			},
			Repository:   model.Repository{
				Url: pkgProfile.RepoUrl,
			},
			Version:      pkgProfile.PkgVersion,
			Dependencies: nil,
			PriFilename:  pkgProfile.PriFilename,
		}
	return rsp, nil
}
