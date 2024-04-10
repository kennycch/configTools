package jwt

import (
	"config_tools/config"
	"config_tools/tools/redis"
	"fmt"
	"time"

	r "github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
	"github.com/kennycch/gotools/general"
)

// GenerateToken 生成Token
func GenerateToken(userId uint32) string {
	tokenData := TokenData{
		Id: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: general.NowAdd(time.Duration(config.Jwt.TimeOut) * time.Hour).UnixMilli(),
			Issuer:    config.Jwt.Issuer,
		},
	}
	// 生成Token
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenData).SignedString([]byte(config.Jwt.SecretKey))
	if err != nil {
		return ""
	}
	return token
}

// ParseToken 解析Token
func ParseToken(tokenStr string) *TokenData {
	// ParseWithClaims 解析token
	tokenClaims, err := jwt.ParseWithClaims(tokenStr, &TokenData{}, func(token *jwt.Token) (interface{}, error) {
		// 使用签名解析用户传入的token,获取载荷部分数据
		return []byte(config.Jwt.SecretKey), nil
	})
	if err != nil {
		return nil
	}
	if tokenClaims != nil {
		//Valid用于校验鉴权声明。解析出载荷部分
		if tokenData, ok := tokenClaims.Claims.(*TokenData); ok &&
			tokenClaims.Valid &&
			tokenData.StandardClaims.Issuer == config.Jwt.Issuer &&
			tokenData.ExpiresAt > general.NowMilli() {
			return tokenData
		}
	}
	return nil
}

// CheckLogout 查询token是否已登出
func CheckLogout(tokenStr string) bool {
	logouts := redis.RD.ZRange(redis.LogoutTokens, 0, -1).Val()
	return general.InArray(logouts, tokenStr)
}

// Logout 登出
func Logout(tokenStr string, expiresAt int64) {
	// Token加入黑名单
	redis.RD.ZAdd(redis.LogoutTokens, r.Z{
		Score:  float64(expiresAt),
		Member: tokenStr,
	}).Result()
	// 删除已过期Token
	now := general.NowMilli()
	redis.RD.ZRemRangeByScore(redis.LogoutTokens, "0", fmt.Sprint(now)).Result()
}
