package controllers

import "github.com/gin-gonic/gin"

//bluebell/controllers/response.go

// ResponseData 响应信息
type ResponseData struct {
	Code ResCode     `json:"code"`           //业务响应状态码
	Msg  interface{} `json:"msg,omitempty"`  //提示信息
	Data interface{} `json:"data,omitempty"` //数据
}

func ResponseError(c *gin.Context, code ResCode) {
	//gin.H{
	//	"cod":"xx",
	//	"msg":"xx",
	//	"data":"xx",
	//}
	rd := &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	c.JSON(200, rd)
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	rd := &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	c.JSON(200, rd)
}
func ResponseSuccess(c *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	}
	c.JSON(200, rd)
}
