package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/showmebug/my-gin-demo/global"
)

type jwtService struct {
}

var JwtService = new(jwtService)

// 所有需要颁发 token 的用户模型必须实现这个接口
type JwtUser interface {
	GetUid() string
}

// CustomClaims 自定义 Claims
type CustomClaims struct {
	jwt.RegisteredClaims
}

const (
	TokenType    = "bearer"
	AppGuardName = "app"
)

type TokenOutPut struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func (jwtService *jwtService) CreateToken(GuardName string, user JwtUser) (tokenData TokenOutPut, err error, token *jwt.Token) {
	token = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		CustomClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.App.Config.Jwt.JwtTtl) * time.Second)),
				ID:        user.GetUid(),
				Issuer:    GuardName,
				NotBefore: jwt.NewNumericDate(time.Now()),
			},
		},
	)
	tokenStr, err := token.SignedString([]byte(global.App.Config.Jwt.Secret))
	tokenData = TokenOutPut{
		tokenStr,
		int(global.App.Config.Jwt.JwtTtl),
		TokenType,
	}
	return
}
