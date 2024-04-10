package controller

import (
	"config_tools/app/dao"
	"config_tools/app/errors"
	"config_tools/app/request"
	"config_tools/app/service"
	"config_tools/tools/net/middleware"

	"github.com/gin-gonic/gin"
)

type tableController struct{}

var (
	tableControllers = tableController{}
)

// TableRoute 路由注册
func TableRoute(route *gin.Engine) {
	// 接口验签
	sRoute := route.Group("", middleware.CheckSign())
	{
		// 需要校验登录的
		lRoute := sRoute.Group("", middleware.CheckLogin())
		{
			// 配置表列表
			lRoute.GET("/table/list", tableControllers.List)
			// 配置表创建
			lRoute.GET("/table/create", tableControllers.Create)
		}
	}
}

// List 配置表列表
func (t *tableController) List(ctx *gin.Context) {
	// 校验参数
	req := &request.TableListRequest{}
	if err := service.CheckValid(ctx, req); err != nil {
		service.JsonResponse(ctx, errors.ErrorCodeParamsError, nil)
		return
	}
	list, count := dao.GetTableList(req)
	service.JsonResponse(ctx, errors.CodeSuccess, gin.H{
		"list":  list,
		"count": count,
	})
}

// Create 配置表创建
func (t *tableController) Create(ctx *gin.Context) {

}
