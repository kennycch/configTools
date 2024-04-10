package service

import (
	"config_tools/app/errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kennycch/gotools/general"
)

func JsonResponse(ctx *gin.Context, code errors.ErrorCode, data interface{}) {
	msg := ""
	if msgData, ok := errors.Msgs[code]; ok {
		msg = msgData
	}
	resMsg := &JsonResponseData{
		Code:    code,
		Message: msg,
		Time:    general.NowUnix(),
		Data:    data,
	}
	ctx.JSON(http.StatusOK, resMsg)
}
