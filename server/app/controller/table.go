package controller

import (
	"config_tools/app/dao"
	"config_tools/app/errors"
	"config_tools/app/request"
	"config_tools/app/service"
	"config_tools/tools/excel"
	"config_tools/tools/git"
	"config_tools/tools/go_struct"
	"config_tools/tools/json"
	"config_tools/tools/net/middleware"
	"fmt"

	"github.com/kennycch/gotools/sort"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
			lRoute.POST("/table/create", tableControllers.Create)
			// 配置表编辑
			lRoute.PUT("/table/edit", tableControllers.Edit)
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
	// 用方法锁让请求线性
	service.LockFunc("Table", func() {
		// 校验参数
		req := &request.TableCreateRequest{}
		if err := service.CheckValid(ctx, req); err != nil {
			service.JsonResponse(ctx, errors.ErrorCodeParamsError, nil)
			return
		}
		// 判断配置表是否已存在
		table := &dao.Table{
			GameId: req.GameId,
			Name:   req.Name,
		}
		dao.GetTableByGameIdAndName(table)
		if table.Id > 0 {
			service.JsonResponse(ctx, errors.ErrorCodeTableAlreadyExist, nil)
			return
		}
		// 游戏是否存在
		game := &dao.Game{
			Model: dao.Model{
				Id: req.GameId,
			},
		}
		dao.GetGameById(game)
		if game.Id == 0 {
			service.JsonResponse(ctx, errors.ErrorCodeGameNotFound, nil)
			return
		}
		// 开启事务
		fields := []*dao.Field{}
		err := dao.DB.Transaction(func(tx *gorm.DB) error {
			// 创建配置表
			table = &dao.Table{
				GameId:    req.GameId,
				Name:      req.Name,
				Comment:   req.Comment,
				Hash:      service.GetPassword(fmt.Sprintf("%d_%s_%d", req.GameId, req.Name, 0)),
				TableType: req.TableType,
				Status:    dao.TableStatusCreated,
			}
			if err := tx.Create(table).Error; err != nil {
				return err
			}
			// 创建字段
			req.Fields = sort.Heap(req.Fields, sort.DESC, func(field request.CreateFieldRequest) int {
				return field.Id
			})
			idMap := map[int]uint32{}
			for _, field := range req.Fields {
				idMap[field.Id] = 0
			}
			for _, f := range req.Fields {
				field := &dao.Field{
					TableId:   table.Id,
					Name:      f.Name,
					Chinese:   f.Chinese,
					Comment:   f.Comment,
					FieldType: f.FieldType,
					Example:   f.Example,
				}
				parentId, ok := idMap[f.ParentId]
				if ok {
					field.ParentId = parentId
				}
				if err := tx.Create(field).Error; err != nil {
					return err
				}
				idMap[f.Id] = field.Id
				fields = append(fields, field)
			}
			return nil
		})
		if err != nil {
			service.JsonResponse(ctx, errors.CodeServerError, nil)
			return
		}
		service.JsonResponse(ctx, errors.CodeSuccess, gin.H{
			"table":  table,
			"fields": fields,
		})
		// 生成Excel
		excelPath := git.GetTargetPath(game.ExcelGit, git.Dev)
		excel.GenerateDemo(table, fields, excelPath)
		excelFullPath := fmt.Sprintf("%s/%s.xlsx", excelPath, table.Comment)
		// 推送git
		git.Push(game.ExcelGit, fmt.Sprintf("创建配置表：%s", table.Comment), git.Dev)
		// 生成json文件
		jsonPath := git.GetTargetPath(game.ClientGit, git.Dev)
		json.StructureJson(excelFullPath, jsonPath)
		git.Push(game.ClientGit, fmt.Sprintf("创建配置表Json：%s", table.Comment), git.Dev)
		// 生成go文件
		goPath := git.GetTargetPath(game.ServerGit, git.Dev)
		go_struct.StructureGo(excelFullPath, goPath)
		git.Push(game.ServerGit, fmt.Sprintf("创建配置表Go：%s", table.Comment), git.Dev)
	})
}

// Edit 配置表编辑
func (t *tableController) Edit(ctx *gin.Context) {
	// 用方法锁让请求线性
	service.LockFunc("Table", func() {
		// 校验参数
		req := &request.TableEditRequest{}
		if err := service.CheckValid(ctx, req); err != nil {
			service.JsonResponse(ctx, errors.ErrorCodeParamsError, nil)
			return
		}
		// 游戏是否存在
		game := &dao.Game{
			Model: dao.Model{
				Id: req.GameId,
			},
		}
		dao.GetGameById(game)
		if game.Id == 0 {
			service.JsonResponse(ctx, errors.ErrorCodeGameNotFound, nil)
			return
		}
		// 获取配置表
		table := &dao.Table{
			Model: dao.Model{
				Id: req.Id,
			},
		}
		if err := dao.GetTableById(table); err != nil {
			service.JsonResponse(ctx, errors.ErrorCodeTableNotFound, nil)
			return
		} else if table.Status == dao.TableStatusPublished { // 已发布配置表不允许修改
			service.JsonResponse(ctx, errors.ErrorCodePublishedTableCanNotEdit, nil)
			return
		}
		// 开启事务
		// fields := []*dao.Field{}
		// oldName, oldComment = table.Name, table.Comment
		err := dao.DB.Transaction(func(tx *gorm.DB) error {
			
			return nil
		})
		if err != nil {
			service.JsonResponse(ctx, errors.CodeServerError, nil)
			return
		}
	})
}
