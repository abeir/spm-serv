package service

import (
	"spm-serv/core"
	"spm-serv/dao"
	"spm-serv/model"
	"spm-serv/model/po"
)

func SearchService(req *model.SearchRequest) (*model.SearchResponse, error){
	core.Log.Debugf("SearchService request: %+v\n", req)

	var pkgProfiles []po.PackageProfile

	pkgProfiles, err := dao.PackageProfileDaoImpl.SelectLastVersionByPkgNameLike("%" + req.PackageName + "%")
	if err!=nil {
		return nil, err
	}
	var data []*model.SearchResponseData
	for _, pkg := range pkgProfiles {
		data = append(data, &model.SearchResponseData{
			Name:        pkg.PkgName,
			Description: pkg.PkgDesc,
		})
	}
	rsp := &model.SearchResponse{
		BaseResponse: model.BaseResponse{
			Code: model.CODE_SUCCESS,
			Msg:  "ok",
		},
		Data: data,
	}
	return rsp, nil
}
