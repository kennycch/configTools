package http

import (
	"config_tools/tools/net/middleware"

	"github.com/gin-gonic/gin"
)

func (h *HttpService) Route(route *gin.Engine) {
	// 接口验签
	route.Use(middleware.CheckSign())
}
