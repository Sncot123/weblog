package controllers

import (
	"fmt"
	"web_app/logic"
	"web_app/models"

	"go.uber.org/zap"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

// PostVoteHandler 帖子投票的接口
// @Summary 用于给帖子投票1
// @Description 给帖子投票的接口2
// @Tags 帖子相关的接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.VoteData false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} controllers.ResponseData
// @Router /vote [post]
func PostVoteHandler(c *gin.Context) {
	p := new(models.VoteData)
	err := c.ShouldBindJSON(p)

	if err != nil {
		errs, ok := err.(validator.ValidationErrors) //类型断言
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, errs)
		return
	}
	userID, err := getCurrentUserID(c)

	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	err = logic.PostVote(userID, p)
	if err != nil {
		zap.L().Debug("logic.PostVote(userID, p)",
			zap.Int64("userID", userID),
			zap.String("postID", p.PostID),
			zap.Int8("direction", p.Direction))
		ResponseError(c, CodeServerBusy)
		return
	}
	fmt.Println("-----------------------")
	ResponseSuccess(c, nil)
}
