package middleware

import (
	"config_tools/app/errors"
	"config_tools/app/service"
	"config_tools/config"
	"config_tools/tools/jwt"
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kennycch/gotools/general"
)

// 登录校验
func CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取token
		authHeader := ctx.Request.Header.Get("Authorization")
		// 校验入参
		if authHeader == "" {
			service.JsonResponse(ctx, errors.CodeTokenInvalid, nil)
			ctx.Abort()
			return
		}
		// 裁剪token
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			service.JsonResponse(ctx, errors.CodeTokenInvalid, nil)
			ctx.Abort()
			return
		}
		// 解析token
		claims := jwt.ParseToken(parts[1])
		if claims == nil ||
			claims.StandardClaims.Issuer != config.Jwt.Issuer ||
			claims.StandardClaims.ExpiresAt < general.NowUnix() {
			service.JsonResponse(ctx, errors.CodeTokenInvalid, nil)
			ctx.Abort()
			return
		}
		// 查看是否已登出
		if jwt.CheckLogout(parts[1]) {
			service.JsonResponse(ctx, errors.CodeTokenInvalid, nil)
			ctx.Abort()
			return
		}
		// 设置用户和Token
		tokenData, _ := json.Marshal(claims)
		ctx.Request.Header.Set("token", parts[1])
		ctx.Request.Header.Set("tokenData", string(tokenData))
		ctx.Next()
	}
}
