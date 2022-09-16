package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"web_app/models"

	"go.uber.org/zap"
)

//把每一步数据库操作封装成一个函数，根据logic根据业务需求调用
const secret = "sncot"

// CheckUserExists 检查指定用户名的用户是否存在
func CheckUserExists(username string) (err error) {
	sqlStr := "SELECT count(user_id) FROM user WHERE username=?"
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser 注册用户
func InsertUser(user *models.User) (err error) {
	//对密码进行加密
	password := encryptPassword(user.Password)
	//执行sql语句入库
	sqlStr := "INSERT INTO user(user_id,username,password) VALUES(?,?,?)"
	_, err = db.Exec(sqlStr, user.UserID, user.UserName, password)
	if err != nil {
		return err
	}
	return
}
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error) {
	oldPassword := encryptPassword(user.Password)
	sqlStr := "SELECT user_id,username,password FROM user Where username = ?"
	err = db.Get(user, sqlStr, user.UserName)
	if err == sql.ErrNoRows {
		zap.L().Error("syntax login failed ", zap.Error(err))
		return ErrorUserNotExist
	}
	if err != nil {
		//查询失败
		return err
	}
	if oldPassword != user.Password {
		zap.L().Error("wrong password ", zap.String(user.UserName, "wrong password"))
		return ErrorInvalidPassword
	}
	return
}
