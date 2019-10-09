package controller

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"spm-serv/model"
	"spm-serv/model/spm"
	"spm-serv/service"
)

var spmService = &service.SpmService{}

type SpmController struct {
}

//回响测试
func (s *SpmController) Ping(c *gin.Context){
	c.JSON(http.StatusOK, model.BaseResponse{
		Code: model.CODE_SUCCESS,
		Msg:  "pong",
	})
}

//发布包
func (s *SpmController) Publish(c *gin.Context) {
	req := &spm.PublishRequest{}
	err := c.ShouldBindJSON(req)
	if err!=nil {
		c.JSON(http.StatusOK, model.FailResponse(req.GetError(err.(validator.ValidationErrors))))
		return
	}
	rsp, err := spmService.PublicshService(req)
	if err!=nil {
		c.JSON(http.StatusOK, model.FailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, rsp)
}

//检索包
func (s *SpmController) Search(c *gin.Context) {
	req := &spm.SearchRequest{}
	err := c.ShouldBindQuery(req)
	if err!=nil {
		c.JSON(http.StatusOK, model.FailResponse(req.GetError(err.(validator.ValidationErrors))))
		return
	}
	rsp, err := spmService.SearchService(req)
	if err!=nil {
		c.JSON(http.StatusOK, model.FailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, rsp)
}

//查询包信息
func (s *SpmController) Info(c *gin.Context) {
	req := &spm.InfoRequest{}
	err := c.ShouldBindQuery(req)
	if err!=nil {
		c.JSON(http.StatusOK, model.FailResponse(req.GetError(err.(validator.ValidationErrors))))
		return
	}
	rsp, err := spmService.InfoService(req)
	if err!=nil {
		c.JSON(http.StatusOK, model.FailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, rsp)
}

//检查spm更新
func (s *SpmController) Upgrade(c *gin.Context){
	req := &spm.UpgradeRequest{}
	err := c.ShouldBindQuery(req)
	if err!=nil {
		c.JSON(http.StatusOK, model.FailResponse(req.GetError(err.(validator.ValidationErrors))))
		return
	}
	rsp, err := spmService.UpgradeService(req)
	if err!=nil {
		c.JSON(http.StatusOK, model.FailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, rsp)
}

//下载spm
//参数version可为latest，表示下载最新版本
func (s *SpmController) Download(c *gin.Context){
	req := &spm.UpgradeRequest{}
	err := c.ShouldBindQuery(req)
	if err!=nil {
		c.JSON(http.StatusOK, model.FailResponse(req.GetError(err.(validator.ValidationErrors))))
		return
	}
	path, err := spmService.DownloadVersion(req)
	if err!=nil {
		c.JSON(http.StatusOK, model.FailResponse(err.Error()))
		return
	}
	if path=="" {
		c.JSON(http.StatusOK, model.FailResponse("File does not exist"))
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment;filename=spm")
	c.File(path)
}