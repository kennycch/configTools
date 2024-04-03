package net

import (
	"config_tools/tools/net/http"
	"config_tools/tools/net/pprof"

	"github.com/gin-gonic/gin"
)

type RouteService interface {
	Route(route *gin.Engine)
}

var (
	// Gin路由对象
	route = &gin.Engine{}
	// 要注册的服务路由
	routeServices = []RouteService{
		&pprof.PprofService{},
		&http.HttpService{},
	}
)
