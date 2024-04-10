package http

import (
	"config_tools/app/controller"

	"github.com/gin-gonic/gin"
)

func (h *HttpService) Route(route *gin.Engine) {
	// 文件相关
	controller.FileRoute(route)
	// 账号相关路由注册
	controller.AccountRoute(route)
	// 游戏相关路由注册
	controller.GameRoute(route)
	// 配置表相关路由注册
	controller.TableRoute(route)
}
