package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const CtxUserID = "user_id"

var ErrorUserNotLogin = errors.New("用户未登录")

// getCurrentUser 获取当前用户id
func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserID)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return userID, err
}
