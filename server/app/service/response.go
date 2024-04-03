package service

import (
	"config_tools/app/errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JsonResponse(ctx *gin.Context, code errors.ErrorCode, data interface{}) {
	msg := ""
	if msgData, ok := errors.Msgs[code]; ok {
		msg = msgData
	}
	resMsg := &JsonResponseData{
		Code:    code,
		Message: msg,
		Data:    data,
		Time:    time.Now().Unix(),
	}
	ctx.JSON(http.StatusOK, resMsg)
}
