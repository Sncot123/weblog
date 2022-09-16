package controllers

import (
	"errors"
	"fmt"
	"web_app/dao/mysql"
	"web_app/logic"
	"web_app/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// SignUpHandler 注册用户的接口
// @Summary 注册用户1
// @Description 注册用户
// @Tags 登录相关的接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.User false "查询参数"
// @Security NoApiKeyAuth
// @Success 200 {object} controllers.ResponseData
// @Router /signup [post]
func SignUpHandler(c *gin.Context) {
	// 1、获取参数和参数校验
	var p = new(models.ParamSignUp)
	err := c.ShouldBindJSON(p) //只能对请求的数据类型和格式进行校对（是否为json格式）
	if err != nil {
		//请求参数有误，直接返回响应
		zap.L().Error("signup with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	////手动对请求数据进行详细的业务规则校验
	//if len(p.UserName) == 0 || len(p.Password) == 0 || p.Password != p.RePassword {
	//	zap.L().Error("signup with invalid param")
	//	c.JSON(200, gin.H{
	//		"msg": "请求参数有误2",
	//	})
	//	return
	//}
	fmt.Println(p)
	// 2、业务处理
	if err := logic.SignUp(p); err != nil {
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3、返回响应
	c.JSON(200, gin.H{"msg": "register successed"})

}

// LoginHandler 登陆的接口
// @Summary 登录1
// @Description 根据账号密码登录的接口
// @Tags 登录相关的接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.User false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} controllers.ResponseData
// @Router /Login [post]
func LoginHandler(c *gin.Context) {
	// 1、获取参数
	user := new(models.ParamLogin)
	err := c.ShouldBindJSON(user)
	if err != nil {
		zap.L().Error("login with invalid param", zap.String("username", user.UserName), zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2、逻辑处理
	var token string
	token, err = logic.Login(user)
	if err != nil {
		zap.L().Error("login failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	// 3、返回响应
	ResponseSuccess(c, token)
	return
}
