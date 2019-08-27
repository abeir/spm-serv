package service

import "spm-serv/model/pkg"

type PackageService struct {
}

func (p *PackageService) List(req *pkg.PackageListReq) (*pkg.PackageListRsp, error){

	return nil, nil
}

func (p *PackageService) Info(req *pkg.PackageInfoReq) (*pkg.PackageInfoRsp, error){

	return nil, nil
}

func (p *PackageService) Enable(req *pkg.PackageStatusReq) (*pkg.PackageStatusRsp, error){
	return nil, nil
}

func (p *PackageService) Disable(req *pkg.PackageStatusReq) (*pkg.PackageInfoRsp, error){
	return nil, nil
}
