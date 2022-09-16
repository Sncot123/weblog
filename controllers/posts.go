package controllers

import (
	"strconv"
	"web_app/logic"
	"web_app/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetPostListHandler 获取帖子列表的接口
// @Summary 获取社区帖子列表信息1
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 社区相关的接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ApiPost false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} controllers.ResponseData
// @Router /posts [get]
func GetPostListHandler(c *gin.Context) {
	// 1、获取数据
	page, size := getPageInfo(c)
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 2、返回响应
	ResponseSuccess(c, data)
}

// GetPostsListHandler2 升级版的获取帖子列表 根据前端传回的参数要求获取指定方式排序的帖子
// @Summary 根据帖子的创建时间或者帖子的分数排序获取得帖子列表的接口1
// @Description 按时间或分数排序查询帖子列表的接口
// @Tags 帖子相关的接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.CommunityDetail false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} controllers.ResponseData
// @Router /post2 [get]
func GetPostsListHandler2(c *gin.Context) {
	// 1、获取参数
	// 2、从redis中获取指定排序的帖子的id
	// 3、根据id去数据库中查询帖子的详细信息
	//Get请求参数、api/v1/posts2?page=1&size=10&order=time
	//c.ShouldBindJSON()  如果请求中携带的是json格式数据，采用此方法
	//c.ShouldBind()  根据请求的数据类型选择相应的方法获取数据
	// 获取请求参数
	p := models.ParamPostsList2{
		Page:  1,
		Size:  10,
		Order: models.OrderTime, //代码中禁止出现 magic string
	}
	err := c.ShouldBindQuery(p)
	if err != nil {
		zap.L().Error("GetPostsListHandler2(c *gin.Context)升级版 with invalid param  获取参数失败")
		ResponseError(c, CodeServerBusy)
		return
	}
	//获取数据
	data, err := logic.GetPostsList2(p)
	if err != nil {
		zap.L().Error("logic.GetPostsList2(page, size) failed!")
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
func getPageInfo(c *gin.Context) (page, size int64) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")
	var (
		err error
	)
	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		zap.L().Error("page:parse failed ")
		return
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		zap.L().Error("size:parse failed ")
		return
	}
	return
}
