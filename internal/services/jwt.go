package services

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/showmebug/my-gin-demo/global"
	"github.com/showmebug/my-gin-demo/internal/pkg"
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

// 获取黑名单缓存 key
func (jwtService *jwtService) getBlackListKey(tokenStr string) string {

	result := "jwt_black_list:" + pkg.MD5([]byte(tokenStr))
	fmt.Println(result)
	return result
	// return "jwt_black_list:" + pkg.MD5([]byte(tokenStr))
}

// JoinBlackList token 加入黑名单
func (jwtService *jwtService) JoinBlackList(token *jwt.Token) (err error) {
	// 获取当前时间的Unix时间戳
	nowUnix := time.Now().Unix()

	// 安全地进行类型断言
	customClaims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return fmt.Errorf("token claims are not of type *CustomClaims")
	}

	// 检查ExpiresAt是否为nil，确保不会解引用nil指针
	if customClaims.ExpiresAt == nil {
		return fmt.Errorf("token has no expiration time")
	}

	// 将*jwt.NumericDate转换为int64类型的Unix时间戳并计算时间差
	expiresAtUnix := customClaims.ExpiresAt.Unix()
	if expiresAtUnix <= nowUnix {
		return fmt.Errorf("token is already expired or expires at the current time")
	}
	timer := time.Duration(expiresAtUnix-nowUnix) * time.Second

	// 将token剩余时间设置为缓存有效期，并将当前时间作为缓存value值
	if err := global.App.Redis.SetNX(context.Background(), jwtService.getBlackListKey(token.Raw), nowUnix, timer).Err(); err != nil {
		// 处理Redis操作错误
		return fmt.Errorf("error setting token to blacklist in Redis: %w", err)
	}

	return nil
}

// func (jwtService *jwtService) JoinBlackList(token *jwt.Token) (err error) {
// 	nowUnix := time.Now().Unix()
// 	timer := time.Duration(token.Claims.(*CustomClaims).ExpiresAt - nowUnix) * time.Second
// 	// 将 token 剩余时间设置为缓存有效期，并将当前时间作为缓存 value 值
// 	err = global.App.Redis.SetNX(context.Background(), jwtService.getBlackListKey(token.Raw), nowUnix, timer).Err()
// 	return
// }

// IsInBlacklist token 是否在黑名单中
func (jwtService *jwtService) IsInBlacklist(tokenStr string) bool {
	joinUnixStr, err := global.App.Redis.Get(context.Background(), jwtService.getBlackListKey(tokenStr)).Result()
	joinUnix, err := strconv.ParseInt(joinUnixStr, 10, 64)
	if joinUnixStr == "" || err != nil {
		return false
	}
	// JwtBlacklistGracePeriod 为黑名单宽限时间，避免并发请求失效
	if time.Now().Unix()-joinUnix < global.App.Config.Jwt.JwtBlacklistGracePeriod {
		return false
	}
	return true
}
