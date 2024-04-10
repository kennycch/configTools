package service

import "github.com/gin-gonic/gin"

func CheckValid(ctx *gin.Context, requset interface{}) error {
	var err error
	// 根据请求头判断是否json传入
	contentType := ctx.Request.Header.Get("Content-Type")
	if contentType == "application/json" {
		err = ctx.ShouldBindJSON(requset)
	} else {
		err = ctx.ShouldBind(requset)
	}
	return err
}
