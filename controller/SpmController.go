package controller

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"spm-serv/model"
	"spm-serv/service"
)

func Ping(c *gin.Context){
	c.JSON(http.StatusOK, model.BaseResponse{
		Code: model.CODE_SUCCESS,
		Msg:  "pong",
	})
}

func Publish(c *gin.Context) {
	req := &model.PublishRequest{}
	err := c.ShouldBindJSON(req)
	if err!=nil {
		c.JSON(http.StatusOK, model.FailResponse(req.GetError(err.(validator.ValidationErrors))))
		return
	}
	rsp, err := service.PublicshService(req)
	if err!=nil {
		c.JSON(http.StatusOK, model.FailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, rsp)
}

func Search(c *gin.Context) {
	req := &model.SearchRequest{}
	err := c.ShouldBindQuery(req)
	if err!=nil {
		c.JSON(http.StatusOK, model.FailResponse(req.GetError(err.(validator.ValidationErrors))))
		return
	}
	rsp, err := service.SearchService(req)
	if err!=nil {
		c.JSON(http.StatusOK, model.FailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, rsp)
}

func Info(c *gin.Context) {
	req := &model.InfoRequest{}
	err := c.ShouldBindQuery(req)
	if err!=nil {
		c.JSON(http.StatusOK, model.FailResponse(req.GetError(err.(validator.ValidationErrors))))
		return
	}
	rsp, err := service.InfoService(req)
	if err!=nil {
		c.JSON(http.StatusOK, model.FailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, rsp)
}

func Upgrade(c *gin.Context){
	req := &model.UpgradeRequest{}
	err := c.ShouldBindQuery(req)
	if err!=nil {
		c.JSON(http.StatusOK, model.FailResponse(req.GetError(err.(validator.ValidationErrors))))
		return
	}
	rsp, err := service.UpgradeService(req)
	if err!=nil {
		c.JSON(http.StatusOK, model.FailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, rsp)
}

func Download(c *gin.Context){
	req := &model.UpgradeRequest{}
	err := c.ShouldBindQuery(req)
	if err!=nil {
		c.JSON(http.StatusOK, model.FailResponse(req.GetError(err.(validator.ValidationErrors))))
		return
	}
	path, err := service.DownloadVersion(req)
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