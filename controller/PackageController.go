package controller

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"spm-serv/model"
	"spm-serv/model/pkg"
	"spm-serv/service"
)

var pkgService = &service.PackageService{}

type PackageController struct {
}

//查询列表
func (p *PackageController) List(ct *gin.Context){
	req := &pkg.PackageListReq{}
	err := ct.ShouldBindJSON(req)
	if err!=nil {
		ct.JSON(http.StatusOK, model.FailResponse(err.Error()))
		return
	}
	rsp, err := pkgService.List(req)
	if err!=nil {
		ct.JSON(http.StatusOK, model.FailResponse(err.Error()))
		return
	}
	ct.JSON(http.StatusOK, rsp)
}

//查询详情
func (p *PackageController) Info(ct *gin.Context){
	req := &pkg.PackageInfoReq{}
	err := ct.ShouldBindJSON(req)
	if err!=nil {
		ct.JSON(http.StatusOK, model.FailResponse(req.GetError(err.(validator.ValidationErrors))))
		return
	}
	rsp, err := pkgService.Info(req)
	if err!=nil {
		ct.JSON(http.StatusOK, model.FailResponse(err.Error()))
		return
	}
	ct.JSON(http.StatusOK, rsp)
}

//设置为可用
func (p *PackageController) Enable(ct *gin.Context){
	req := &pkg.PackageStatusReq{}
	err := ct.ShouldBindJSON(req)
	if err!=nil {
		ct.JSON(http.StatusOK, model.FailResponse(req.GetError(err.(validator.ValidationErrors))))
		return
	}
	rsp, err := pkgService.Enable(req)
	if err!=nil {
		ct.JSON(http.StatusOK, model.FailResponse(err.Error()))
		return
	}
	ct.JSON(http.StatusOK, rsp)
}

//设置为不可用
func (p *PackageController) Disable(ct *gin.Context){
	req := &pkg.PackageStatusReq{}
	err := ct.ShouldBindJSON(req)
	if err!=nil {
		ct.JSON(http.StatusOK, model.FailResponse(req.GetError(err.(validator.ValidationErrors))))
		return
	}
	rsp, err := pkgService.Disable(req)
	if err!=nil {
		ct.JSON(http.StatusOK, model.FailResponse(err.Error()))
		return
	}
	ct.JSON(http.StatusOK, rsp)
}
