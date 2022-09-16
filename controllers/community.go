package controllers

import (
	"strconv"
	"web_app/logic"
	"web_app/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//-------和社区相关的---------

//CommunityHandler 查询社区信息接口
// @Summary 查询社区的接口1
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 社区相关的接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.Community false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} controllers.ResponseData
// @Router /community [get]
func CommunityHandler(c *gin.Context) {
	//查询到所有社区的id、name以列表形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) //不轻易把服务器端错误暴露给前端
		return
	}
	ResponseSuccess(c, data)
}

// CommunityDetailHandler 获取社区详细信息的接口
// @Summary 获取社区详细信息的接口1
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 社区相关的接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.CommunityDetail false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} controllers.ResponseData
// @Router /community/:id [get]
func CommunityDetailHandler(c *gin.Context) {
	//1、获取社区id
	idSTR := c.Param("id")
	id, err := strconv.ParseInt(idSTR, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	//2、查询到所有的社区，以列表的形式返回
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

//GetCommunityPostsListHandler 按社区分类 根据帖子的创建时间或者帖子分数的顺序来获取帖子的排序
// @Summary 获取社区帖子列表信息1
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 社区相关的接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.CommunityDetail false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} controllers.ResponseData
// @Router /postByCommunity [get]
func GetCommunityPostsListHandler(c *gin.Context) {
	//Get请求参数、api/v1/posts2?page=1&size=10&order=time
	//c.ShouldBindJSON()  如果请求中携带的是json格式数据，采用此方法
	//c.ShouldBind()  根据请求的数据类型选择相应的方法获取数据
	// 获取请求参数
	p := models.ParamPostsListCommunity{
		ParamPostsList2: &models.ParamPostsList2{
			Page:  1,
			Size:  10,
			Order: models.OrderTime, //代码中禁止出现 magic string},
		},
		CommunityId: 1,
	}
	err := c.ShouldBindQuery(p)
	if err != nil {
		zap.L().Error("GetCommunityPostsListHandler(c *gin.Context) with invalid param  获取参数失败")
		ResponseError(c, CodeServerBusy)
		return
	}
	//获取数据
	data, err := logic.GetCommunityPostsList(p)
	if err != nil {
		zap.L().Error("logic.GetPostsList2(page, size) failed!")
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
