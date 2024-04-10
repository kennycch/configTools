package service

import (
	"config_tools/app/dao"
	"config_tools/config"
	"config_tools/tools/jwt"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/kennycch/gotools/general"
)

// 生成密码
func GetPassword(password string) string {
	password = fmt.Sprintf("%s._%s_.%s",
		config.Sign.SignKey,
		password,
		config.Jwt.SecretKey,
	)
	return general.Md5(password)
}

// GetToken 获取Token
func GetToken(ctx *gin.Context) string {
	return ctx.Request.Header.Get("token")
}

// GetTokenData 获取账号
func GetTokenData(ctx *gin.Context) *jwt.TokenData {
	tokenData := &jwt.TokenData{}
	str := ctx.Request.Header.Get("tokenData")
	json.Unmarshal([]byte(str), tokenData)
	return tokenData
}

// GetAccount 获取账号
func GetAccount(ctx *gin.Context) *dao.Account {
	tokenData := GetTokenData(ctx)
	account := &dao.Account{
		Model: dao.Model{
			Id: tokenData.Id,
		},
	}
	dao.GetAccountById(account)
	return account
}
