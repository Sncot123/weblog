package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/jwt"
	"web_app/pkg/snowflake"
)

//存放业务逻辑的代码

func SignUp(p *models.ParamSignUp) (err error) {
	// 1、判断用户是否存在
	err = mysql.CheckUserExists(p.UserName)
	if err != nil {
		return err
	}
	// 2、生成uid
	userID := snowflake.GenID()
	// 创建一个user实例
	user := &models.User{
		UserID:   userID,
		UserName: p.UserName,
		Password: p.Password,
	}
	// 3、存入数据库
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		UserName: p.UserName,
		Password: p.Password,
	}
	err = mysql.Login(user)
	if err != nil {
		return "", err
	}
	token, err = jwt.GenToken(user.UserID, user.UserName)
	return token, err
}
