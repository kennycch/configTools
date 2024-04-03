package middleware

import (
	"config_tools/app/errors"
	"config_tools/app/service"
	"config_tools/config"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kennycch/gotools/general"
)

func CheckSign() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取签名三要素
		timeStamp := ctx.Request.Header.Get("timeStamp")
		randomStr := ctx.Request.Header.Get("randomStr")
		signature := ctx.Request.Header.Get("signature")
		if timeStamp == "" || randomStr == "" || signature == "" {
			service.JsonResponse(ctx, errors.ErrorCodeApiAuthFail, nil)
			ctx.Abort()
			return
		}
		// 对比请求时间与服务器时间
		timeStampNum, err := strconv.ParseInt(timeStamp, 10, 64)
		if err != nil {
			service.JsonResponse(ctx, errors.ErrorCodeApiAuthFail, nil)
			ctx.Abort()
			return
		}
		now := time.Now().Unix()
		if int64(math.Abs(float64(now-timeStampNum))) > config.Sign.TimeOut {
			service.JsonResponse(ctx, errors.ErrorCodeApiAuthFail, nil)
			ctx.Abort()
			return
		}
		// 生成签名作对比
		serverSign := getSign(timeStamp, randomStr)
		if serverSign != signature {
			service.JsonResponse(ctx, errors.ErrorCodeApiAuthFail, nil)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

// 生成签名
func getSign(timestamp string, randomStr string) string {
	// 拼接字符串
	str := timestamp + randomStr + config.Sign.SignKey
	// sha1加密
	str = general.Sha1(str)
	// MD5加密
	str = general.Md5(str)
	// 转化成大写
	str = strings.ToUpper(str)
	return str
}
