package controller

import (
	"config_tools/app/dao"
	"config_tools/app/errors"
	"config_tools/app/request"
	"config_tools/app/service"
	"config_tools/tools/net/middleware"

	"github.com/gin-gonic/gin"
)

type gameController struct{}

var (
	gameControllers = gameController{}
)

// GameRoute 路由注册
func GameRoute(route *gin.Engine) {
	// 接口验签
	sRoute := route.Group("", middleware.CheckSign())
	{
		// 需要校验登录的
		lRoute := sRoute.Group("", middleware.CheckLogin())
		{
			// 游戏列表
			lRoute.GET("/game/list", gameControllers.List)
			// 游戏新建
			lRoute.POST("/game/create", gameControllers.Create)
			// 游戏编辑
			lRoute.PUT("/game/edit", gameControllers.Edit)
		}
	}
}

// List 游戏列表
func (g *gameController) List(ctx *gin.Context) {
	// 校验参数
	req := &request.ListBaseRequest{}
	if err := service.CheckValid(ctx, req); err != nil {
		service.JsonResponse(ctx, errors.ErrorCodeParamsError, nil)
		return
	}
	list, count := dao.GetGameList(req)
	service.JsonResponse(ctx, errors.CodeSuccess, gin.H{
		"list":  list,
		"count": count,
	})
}

// Create 游戏新建
func (g *gameController) Create(ctx *gin.Context) {
	// 用方法锁让请求线性
	service.LockFunc("GameCreate", func() {
		// 校验参数
		req := &request.GameCreateRequest{}
		if err := service.CheckValid(ctx, req); err != nil {
			service.JsonResponse(ctx, errors.ErrorCodeParamsError, nil)
			return
		}
		// 查看游戏是否已存在
		oldGame := &dao.Game{
			Name: req.Name,
		}
		dao.GetGameByName(oldGame)
		if oldGame.Id > 0 {
			service.JsonResponse(ctx, errors.ErrorCodeGameAlreadyExist, nil)
			return
		}
		// 创建游戏
		newGame := &dao.Game{
			Name:       req.Name,
			Background: req.Background,
		}
		dao.CreateGame(newGame)
		service.JsonResponse(ctx, errors.CodeSuccess, nil)
	})
}

// Edit 游戏编辑
func (g *gameController) Edit(ctx *gin.Context) {
	// 校验参数
	req := &request.GameEditRequest{}
	if err := service.CheckValid(ctx, req); err != nil {
		service.JsonResponse(ctx, errors.ErrorCodeParamsError, nil)
		return
	}
	// 更新游戏
	dao.GameUpdate(req.Id, map[string]interface{}{
		"name":       req.Name,
		"background": req.Background,
	})
	service.JsonResponse(ctx, errors.CodeSuccess, nil)
}
