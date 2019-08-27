package service

import (
	"spm-serv/core"
	"spm-serv/dao"
	"spm-serv/model"
	"spm-serv/model/po"
	"spm-serv/model/spm"
	"time"
)

type SpmService struct {
}


func (s *SpmService) InfoService(req *spm.InfoRequest) (*spm.InfoResponse, error){
	core.Log.Debugf("InfoService request: %+v\n", *req)

	pkgVersion := req.Version
	if req.Version=="" {
		lastVersion := dao.LastVersionDaoImpl.SelectByPkgName(req.PackageName)
		pkgVersion = lastVersion.PkgVersion
	}

	pkgProfile := dao.PackageProfileDaoImpl.SelectByPkgNameAndPkgVersion(req.PackageName, pkgVersion)
	rsp := &spm.InfoResponse{
		BaseResponse: model.BaseResponse{
			Code: model.CODE_SUCCESS,
			Msg: "ok",
		},
	}
	if pkgProfile.IsEmpty() {
		return rsp, nil
	}

	rsp.Data = &spm.InfoResponseData{
		Package:      spm.Package{
			Name: pkgProfile.PkgName,
			Description: pkgProfile.PkgDesc,
		},
		Author:       spm.Author{
			Name: pkgProfile.AuthorName,
			Email: pkgProfile.AuthorEmail,
			Description: pkgProfile.AuthorDesc,
		},
		Repository:   spm.Repository{
			Url: pkgProfile.RepoUrl,
		},
		Version:      pkgProfile.PkgVersion,
		Dependencies: nil,
		PriFilename:  pkgProfile.PriFilename,
	}
	return rsp, nil
}

func (s *SpmService) SearchService(req *spm.SearchRequest) (*spm.SearchResponse, error){
	core.Log.Debugf("SearchService request: %+v\n", req)

	pkgProfiles := dao.PackageProfileDaoImpl.SelectLastVersionByPkgNameLike("%" + req.PackageName + "%")
	var data []*spm.SearchResponseData
	for _, pkg := range pkgProfiles {
		data = append(data, &spm.SearchResponseData{
			Name:        pkg.PkgName,
			Description: pkg.PkgDesc,
		})
	}
	rsp := &spm.SearchResponse{
		BaseResponse: model.BaseResponse{
			Code: model.CODE_SUCCESS,
			Msg:  "ok",
		},
		Data: data,
	}
	return rsp, nil
}

func (s *SpmService) PublicshService(request *spm.PublishRequest) (*spm.PublishResponse, error){
	core.Log.Debugf("PublicshService request: %+v\n", request)

	newPkgProfile := po.PackageProfile{
		Id:			 core.UUID(),
		PkgName:     request.Package.Name,
		PkgDesc:     request.Package.Description,
		RepoUrl:     request.Repository.Url,
		PkgVersion:  request.Version,
		AuthorName:  request.Author.Name,
		AuthorEmail: request.Author.Email,
		AuthorDesc:  request.Author.Description,
		PriFilename: request.PriFilename,
		CreatedAt:	 time.Now(),
		UpdatedAt:	 time.Now(),
	}
	rsp := &spm.PublishResponse{}

	isExists := s.isPackageExists(request.Package.Name, request.Version)
	if isExists {
		rsp.Code = model.CODE_ERROR
		rsp.Msg = "The " + request.Version + " version of the package already exists"
		return rsp, nil
	}
	err := dao.PackageProfileDaoImpl.Tx(func()error{
		dao.PackageProfileDaoImpl.Insert(newPkgProfile)
		err := s.updateLastPackage(&newPkgProfile)
		if err!=nil {
			return err
		}
		return nil
	})
	if err!=nil {
		return nil, err
	}
	rsp.Code = model.CODE_SUCCESS
	rsp.Msg = "ok"
	return rsp, nil
}

func (s *SpmService) isPackageExists(pkgName, version string) bool{
	count := dao.PackageProfileDaoImpl.CountByPkgNameAndPkgVersion(pkgName, version)
	return count > 0
}

//更新最新版本
// 查询last_version，为空时直接插入上传的包记录；
// 不为空时比较版本号，若上传的版本号大于last_version中版本号则更新
func (s *SpmService) updateLastPackage(pkgProfile *po.PackageProfile) error{
	lastVersion := dao.LastVersionDaoImpl.SelectByPkgName(pkgProfile.PkgName)
	if lastVersion.IsEmpty() {
		newLastVersion := po.LastVersion{
			Id: core.UUID(),
			PkgName: pkgProfile.PkgName,
			PkgVersion: pkgProfile.PkgVersion,
			PkgProfileId: pkgProfile.Id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		dao.LastVersionDaoImpl.Insert(newLastVersion)
		return nil
	}else {
		oldVer, err := core.NewVersion(lastVersion.PkgVersion)
		if err!=nil {
			return err
		}
		newVer, err := core.NewVersion(pkgProfile.PkgVersion)
		if err!=nil {
			return err
		}
		compare := newVer.Compare(oldVer)
		if compare > 0 {
			newLastVersion := po.LastVersion{
				PkgName: 		pkgProfile.PkgName,
				PkgVersion: 	pkgProfile.PkgVersion,
				PkgProfileId: 	pkgProfile.Id,
				UpdatedAt: 		time.Now(),
			}
			dao.LastVersionDaoImpl.UpdateByPkgName(newLastVersion)
		}
		return nil
	}
}


//查询新版本
func (s *SpmService) UpgradeService(req *spm.UpgradeRequest) (*spm.UpgradeResponse, error){
	core.Log.Debugf("UpgradeService request: %+v\n", req)

	upgradeVersion := dao.UpgradeVersionDaoImpl.SelectLatestVersion()

	core.Log.Debugf("upgradeVersion %+v\n", upgradeVersion)

	rsp := &spm.UpgradeResponse{}
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
		rsp.Data = &spm.UpgradeResponseData{Version: upgradeVersion.Version}
		return rsp, nil
	}

	rsp.Code = model.CODE_ERROR
	rsp.Msg = "It's the latest version"
	return rsp, nil
}

//下载版本，返回待下载的文件路径
func (s *SpmService) DownloadVersion(req *spm.UpgradeRequest) (string, error){
	core.Log.Debugf("DownloadVersion request: %+v\n", req)

	upgradeVersion := dao.UpgradeVersionDaoImpl.SelectByVersion(req.Version)
	core.Log.Debugf("upgradeVersion: %+v\n", upgradeVersion)
	return upgradeVersion.Path, nil
}
