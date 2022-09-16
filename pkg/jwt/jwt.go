package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

const TokenExpireDuriation = time.Hour * 24

var mySercet = []byte("sncot")

//MyClaims自定义声明结构体并内嵌jwt.StandardClaims
//jwt包自带的jwt.StandardClaims只包含的官方字段
//如果要添加其他字段需要我们自定义结构体
type MyClaims struct {
	UserID   int64
	UserName string
	*jwt.StandardClaims
}

// GenToken 获取token
func GenToken(userID int64, username string) (string, error) {
	//创建一个自己的声明
	c := MyClaims{
		userID,
		username,
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuriation).Unix(), //过期时间
			Issuer:    "bluebell",                                  //签发人
		},
	}
	//使用指定的方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(mySercet)
}

// ParseToken 解析token
func ParseToken(tokenString string) (*MyClaims, error) {
	//解析token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return mySercet, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { //校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
