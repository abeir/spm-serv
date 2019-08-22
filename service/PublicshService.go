package service

import (
	"github.com/abeir/GoMybatis"
	"spm-serv/core"
	"spm-serv/dao"
	"spm-serv/model"
	"spm-serv/model/po"
	"time"
)


func PublicshService(request *model.PublishRequest) (*model.PublishResponse, error){
	core.Log.Debugf("PublicshService request: %+v\n", request)

	newPkgProfile := po.PackageProfile{
		Id:			 dao.UUID(),
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
	rsp := &model.PublishResponse{}

	isExists, err := isPackageExists(request.Package.Name, request.Version)
	if err!=nil {
		return nil, err
	}
	if isExists {
		rsp.Code = model.CODE_ERROR
		rsp.Msg = "The " + request.Version + " version of the package already exists"
		return rsp, nil
	}

	session, err := dao.PackageProfileDaoImpl.NewSession()
	if err!=nil {
		return nil, err
	}
	err = dao.Tx(&session, func()error{
		_, err = dao.PackageProfileDaoImpl.Insert(&session, newPkgProfile)
		if err!=nil {
			return err
		}
		err = updateLastPackage(&session, &newPkgProfile)
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

func isPackageExists(pkgName, version string) (bool, error){
	count, err := dao.PackageProfileDaoImpl.CountByPkgNameAndPkgVersion(pkgName, version)
	if err!=nil {
		return false, err
	}
	return count > 0, nil
}

//更新最新版本
// 查询last_version，为空时直接插入上传的包记录；
// 不为空时比较版本号，若上传的版本号大于last_version中版本号则更新
func updateLastPackage(session *GoMybatis.Session, pkgProfile *po.PackageProfile) error{
	lastVersion, err := dao.LastVersionDaoImpl.SelectByPkgName(pkgProfile.PkgName)
	if err!=nil {
		return err
	}
	if lastVersion.IsEmpty() {
		newLastVersion := &po.LastVersion{
			Id: dao.UUID(),
			PkgName: pkgProfile.PkgName,
			PkgVersion: pkgProfile.PkgVersion,
			PkgProfileId: pkgProfile.Id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		_, err := dao.LastVersionDaoImpl.Insert(session, *newLastVersion)
		if err!=nil {
			return err
		}
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
			newLastVersion := &po.LastVersion{
				PkgName: 		pkgProfile.PkgName,
				PkgVersion: 	pkgProfile.PkgVersion,
				PkgProfileId: 	pkgProfile.Id,
				UpdatedAt: 		time.Now(),
			}
			_, err := dao.LastVersionDaoImpl.UpdateByPkgName(session, *newLastVersion)
			if err!=nil {
				return err
			}
		}
		return nil
	}
}

