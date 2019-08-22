package model

//返回码
const (
	CODE_SUCCESS string = "SUCCESS"
	CODE_ERROR string = "ERROR"
)


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
