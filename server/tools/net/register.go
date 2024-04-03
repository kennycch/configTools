package net

import (
	"config_tools/config"
	"config_tools/tools/lifecycle"
	"config_tools/tools/net/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kennycch/gotools/worker"
)

type NetRegister struct{}

func (n *NetRegister) Start() {
	route = gin.Default()
	// 允许跨域
	route.Use(middleware.Cors())
	// 注册路由
	for _, service := range routeServices {
		service.Route(route)
	}
	// 开启服务
	worker.AddTask(func() {
		route.Run(fmt.Sprintf(":%d", config.Http.Port))
	})
}

func (n *NetRegister) Priority() uint32 {
	return lifecycle.LowPriority
}

func (n *NetRegister) Stop() {

}

func NewNetRegister() *NetRegister {
	return &NetRegister{}
}
