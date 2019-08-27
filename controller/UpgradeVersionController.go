package controller

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"spm-serv/model"
	"spm-serv/model/upgrade"
	"spm-serv/service"
)

var upgradeVersionService = &service.UpgradeVersionService{}

type UpgradeVersionController struct {
}

//查询列表
func (c *UpgradeVersionController) List(ct *gin.Context){
	req := &upgrade.UpgradeVersionListReq{}
	err := ct.ShouldBindJSON(req)
	if err!=nil {
		ct.JSON(http.StatusOK, model.FailResponse(err.Error()))
		return
	}
	rsp, err := upgradeVersionService.List(req)
	if err!=nil {
		ct.JSON(http.StatusOK, model.FailResponse(err.Error()))
		return
	}
	ct.JSON(http.StatusOK, rsp)
}

//查询详情
func (c *UpgradeVersionController) Info(ct *gin.Context){
	req := &upgrade.UpgradeVersionInfoReq{}
	err := ct.ShouldBindJSON(req)
	if err!=nil {
		ct.JSON(http.StatusOK, model.FailResponse(req.GetError(err.(validator.ValidationErrors))))
		return
	}
	rsp, err := upgradeVersionService.Info(req)
	if err!=nil {
		ct.JSON(http.StatusOK, model.FailResponse(err.Error()))
		return
	}
	ct.JSON(http.StatusOK, rsp)
}

//上传spm
func (c *UpgradeVersionController) Upload(ct *gin.Context){
	req := &upgrade.UpgradeVersionUploadReq{}
	err := ct.ShouldBindJSON(req)
	if err!=nil {
		ct.JSON(http.StatusOK, model.FailResponse(req.GetError(err.(validator.ValidationErrors))))
		return
	}
	rsp, err := upgradeVersionService.Upload(req)
	if err!=nil {
		ct.JSON(http.StatusOK, model.FailResponse(err.Error()))
		return
	}
	ct.JSON(http.StatusOK, rsp)
}

//发布spm
func (c *UpgradeVersionController) Release(ct *gin.Context){
	req := &upgrade.UpgradeVersionStatusReq{}
	err := ct.ShouldBindJSON(req)
	if err!=nil {
		ct.JSON(http.StatusOK, model.FailResponse(req.GetError(err.(validator.ValidationErrors))))
		return
	}
	rsp, err := upgradeVersionService.Release(req)
	if err!=nil {
		ct.JSON(http.StatusOK, model.FailResponse(err.Error()))
		return
	}
	ct.JSON(http.StatusOK, rsp)
}

//下架spm
func (c *UpgradeVersionController) Detain(ct *gin.Context){
	req := &upgrade.UpgradeVersionStatusReq{}
	err := ct.ShouldBindJSON(req)
	if err!=nil {
		ct.JSON(http.StatusOK, model.FailResponse(req.GetError(err.(validator.ValidationErrors))))
		return
	}
	rsp, err := upgradeVersionService.Detain(req)
	if err!=nil {
		ct.JSON(http.StatusOK, model.FailResponse(err.Error()))
		return
	}
	ct.JSON(http.StatusOK, rsp)
}