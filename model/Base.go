package model

import "gopkg.in/go-playground/validator.v8"

//返回码
const (
	CODE_SUCCESS string = "SUCCESS"
	CODE_ERROR string = "ERROR"
)


type BaseRequest struct {
	Token string `json:"token"`
}

// 分页参数
type Page struct {
	Page int32 			`json:"page"`
	Row int32 			`json:"row"`
}

func (p *Page) GetPage() int32{
	if p.Page > 0 {
		return p.Page-1
	}
	return 0
}

func (p *Page) GetRow() int32{
	if p.Row > 0 {
		return p.Row
	}
	return 20
}

//BaseResponse 响应返回通用信息
type BaseResponse struct {
	Code string		`json:"code"`
	Msg string		`json:"msg"`
}

func FailResponse(msg string) *BaseResponse{
	return &BaseResponse{
		Code: CODE_ERROR,
		Msg:  msg,
	}
}


type BaseRequester interface {
	GetError(err validator.ValidationErrors) string
}