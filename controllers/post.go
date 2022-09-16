package controllers

import (
	"fmt"
	"strconv"
	"web_app/logic"
	"web_app/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// CreatePostHandler 创建帖子的接口
// @Summary 创建帖子1
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关的接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.Post false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} controllers.ResponseData
// @Router /post [post]
func CreatePostHandler(c *gin.Context) {
	// 1、获取参数及参数的校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(p) failed!  error:", zap.Any("err", err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2、创建帖子
	//从c 取得当前用户的id
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorId = userID
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3、返回响应
	ResponseSuccess(c, nil)
}

// GetPostDetailHandler 获取帖子详细信息的接口
// @Summary 获取帖子详细信息1
// @Description 获取帖子的详细信息2
// @Tags 帖子相关的接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.CommunityDetail false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} controllers.ResponseData
// @Router /post/:id [get]
func GetPostDetailHandler(c *gin.Context) {
	// 1、获取参数 （获取帖子id）
	pid := c.Param("id")
	id, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		zap.L().Error("GetPostDetailHandler: wrong param", zap.Error(err))
		return
	}
	// 2、根据id向数据库查询帖子内容
	data, err := logic.GetPostById(id)
	fmt.Println(err)
	if err != nil {
		zap.L().Error("logic.GetPostById(id) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3、返回响应
	ResponseSuccess(c, data)

}
