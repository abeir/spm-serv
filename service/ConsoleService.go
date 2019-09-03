package service

import (
	"crypto/md5"
	"encoding/hex"
	"spm-serv/controller/middleware"
	"spm-serv/core"
	"spm-serv/dao"
	"spm-serv/model"
	"spm-serv/model/console"
	"strings"
)

type ConsoleService struct {
}

func (c *ConsoleService) Login(req *console.LoginReq) (*console.LoginRsp, error){

	core.Log.Debugf("Login request: %+v\n", *req)

	consoleUser := dao.ConsoleUserDaoImpl.SelectByUserName(req.Username)

	core.Log.Debugf("consoleUser %+v\n", consoleUser)
	rsp := &console.LoginRsp{}
	//数据库中没有记录该用户信息
	if consoleUser.IsEmpty() {
		rsp.Code = model.CODE_ERROR
		rsp.Msg = "No user was found"
		return rsp, nil
	}

	//密码+id
	if strings.ToLower(MD5([]byte(req.Password + consoleUser.Id))) != strings.ToLower(consoleUser.Password){
		rsp.Code = model.CODE_ERROR
		rsp.Msg = "Logon password error"
		return rsp, nil
	}
	//生成token
	tokenString := middleware.SetToken(consoleUser)

	rsp.Code = model.CODE_SUCCESS
	rsp.Msg = "There is a new token"
	rsp.Data = &console.LoginRspData{Token:tokenString}

	core.Log.Debugf("tokenString %+v\n", tokenString)

	return rsp, nil
}

func (c *ConsoleService) Logout(req *console.LogoutReq) (*console.LogoutRsp, error) {
	core.Log.Debugf("Logout request: %+v\n", *req)
	rsp := &console.LogoutRsp{}
	//TODO 清除token

	rsp.Code = model.CODE_SUCCESS
	rsp.Msg = "There is successful exit"

	return rsp, nil
}
//计算MD5签名
func MD5(data []byte) string {
	_md5 := md5.New()
	_md5.Write(data)
	return hex.EncodeToString(_md5.Sum([]byte("")))
}



