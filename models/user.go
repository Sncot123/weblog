package models

//bluebell/models/user.go

// User 用户信息
type User struct {
	UserID   int64  `db:"user_id"`  //用户id
	UserName string `db:"username"` //用户名称
	Password string `db:"password"` //用户密码
}
