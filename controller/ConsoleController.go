package controller

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"spm-serv/model"
	"spm-serv/model/console"
	"spm-serv/service"
)

var consoleService = &service.ConsoleService{}

type ConsoleController struct {
}

func (c *ConsoleController) Login(ct *gin.Context){
	req := &console.LoginReq{}
	err := ct.ShouldBindJSON(req)
	if err!=nil {
		ct.JSON(http.StatusOK, model.FailResponse(req.GetError(err.(validator.ValidationErrors))))
		return
	}
	rsp, err := consoleService.Login(req)
	if err!=nil {
		ct.JSON(http.StatusOK, model.FailResponse(err.Error()))
		return
	}
	ct.JSON(http.StatusOK, rsp)
}

func (c *ConsoleController) Logout(ct *gin.Context){
	req := &console.LogoutReq{}
	err := ct.ShouldBindJSON(req)
	if err!=nil {
		ct.JSON(http.StatusOK, model.FailResponse(err.Error()))
		return
	}
	rsp, err := consoleService.Logout(req)
	if err!=nil {
		ct.JSON(http.StatusOK, model.FailResponse(err.Error()))
		return
	}
	ct.JSON(http.StatusOK, rsp)
}
