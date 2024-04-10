package controller

import (
	"config_tools/app/dao"
	"config_tools/app/errors"
	"config_tools/app/request"
	"config_tools/app/service"
	"config_tools/tools/captcha"
	"config_tools/tools/jwt"
	"config_tools/tools/net/middleware"

	"github.com/gin-gonic/gin"
)

type accountController struct{}

var (
	accountControllers = accountController{}
)

// AccountRoute 路由注册
func AccountRoute(route *gin.Engine) {
	// 接口验签
	sRoute := route.Group("", middleware.CheckSign())
	{
		// 登录验证码
		sRoute.GET("/captcha", accountControllers.Captcha)
		// 登录
		sRoute.POST("/login", accountControllers.Login)
		// 需要校验登录的
		lRoute := sRoute.Group("", middleware.CheckLogin())
		{
			// 修改密码
			lRoute.PUT("/changePassword", accountControllers.ChangePassword)
			// 登出
			lRoute.PUT("/logout", accountControllers.Logout)
			// 账号列表
			lRoute.GET("/account/list", accountControllers.List)
			// 账号创建
			lRoute.POST("/account/create", accountControllers.Create)
			// 账号激活
			lRoute.PUT("/account/activate", accountControllers.Activate)
			// 账号禁用
			lRoute.PUT("/account/deactivate", accountControllers.Deactivate)
			// 账号重置密码
			lRoute.PUT("/account/resetPassword")
		}
	}
}

// Captcha 登录验证码
func (a *accountController) Captcha(ctx *gin.Context) {
	code, img, _, err := captcha.CreateCode()
	if err != nil {
		service.JsonResponse(ctx, errors.CodeServerError, nil)
		return
	}
	service.JsonResponse(ctx, errors.CodeSuccess, gin.H{
		"img":  img,
		"code": code,
	})
}

// Login 登录
func (a *accountController) Login(ctx *gin.Context) {
	// 校验参数
	req := &request.LoginRequest{}
	if err := service.CheckValid(ctx, req); err != nil {
		service.JsonResponse(ctx, errors.ErrorCodeParamsError, nil)
		return
	}
	// 验证码校验
	if !captcha.VerifyCaptcha(req.Code, req.Captcha) {
		service.JsonResponse(ctx, errors.ErrorCodeCaptchaInvalid, nil)
		return
	}
	// 验证账号密码
	account := &dao.Account{
		Account:  req.Account,
		Password: service.GetPassword(req.Password),
	}
	if err := dao.GetAccountByAccountPassword(account); err != nil {
		service.JsonResponse(ctx, errors.ErrorCodeAccountOrPasswordInvalid, nil)
		return
	} else if account.IsActivate == 0 {
		service.JsonResponse(ctx, errors.ErrorCodeAccountNotActivate, nil)
		return
	}
	// 生成Token
	token := jwt.GenerateToken(account.Id)
	if token == "" {
		service.JsonResponse(ctx, errors.CodeServerError, nil)
		return
	}
	service.JsonResponse(ctx, errors.CodeSuccess, gin.H{
		"token":    token,
		"id":       account.Id,
		"account":  account.Account,
		"isSupper": account.IsSupper,
	})
}

// ChangePassword 修改密码
func (a *accountController) ChangePassword(ctx *gin.Context) {
	// 校验参数
	req := &request.ChangePassword{}
	if err := service.CheckValid(ctx, req); err != nil {
		service.JsonResponse(ctx, errors.ErrorCodeParamsError, nil)
		return
	} else if req.NewPassword != req.ConfirmPassword {
		service.JsonResponse(ctx, errors.ErrorCodePasswordsNotMatch, nil)
		return
	}
	// 获取账号信息
	tokenData := service.GetTokenData(ctx)
	account := &dao.Account{
		Model: dao.Model{
			Id: tokenData.Id,
		},
	}
	dao.GetAccountById(account)
	if account.Password != service.GetPassword(req.OldPassword) {
		service.JsonResponse(ctx, errors.ErrorCodeOldPasswordInvalid, nil)
		return
	}
	// 修改密码
	dao.AccountUpdate(tokenData.Id, map[string]interface{}{
		"password": service.GetPassword(req.NewPassword),
	})
	service.JsonResponse(ctx, errors.CodeSuccess, nil)
}

// Logout 登出
func (a *accountController) Logout(ctx *gin.Context) {
	// 获取token和token内容
	token := service.GetToken(ctx)
	tokenData := service.GetTokenData(ctx)
	// 登出
	jwt.Logout(token, tokenData.ExpiresAt)
	service.JsonResponse(ctx, errors.CodeSuccess, nil)
}

// List 账号列表
func (a *accountController) List(ctx *gin.Context) {
	// 校验参数
	req := &request.AccountListRequest{}
	if err := service.CheckValid(ctx, req); err != nil {
		service.JsonResponse(ctx, errors.ErrorCodeParamsError, nil)
		return
	}
	list, count := dao.GetAccountList(req)
	service.JsonResponse(ctx, errors.CodeSuccess, gin.H{
		"list":  list,
		"count": count,
	})
}

// Create 账号创建
func (a *accountController) Create(ctx *gin.Context) {
	// 用方法锁让请求线性
	service.LockFunc("AccountCreate", func() {
		// 校验参数
		req := &request.AccountCreateRequest{}
		if err := service.CheckValid(ctx, req); err != nil {
			service.JsonResponse(ctx, errors.ErrorCodeParamsError, nil)
			return
		}
		// 校验是否超级管理员账号
		account := service.GetAccount(ctx)
		if account.IsSupper == 0 {
			service.JsonResponse(ctx, errors.ErrorCodeHasNotPermission, nil)
			return
		}
		// 查看账号是否已存在
		oldAccount := &dao.Account{
			Account: req.Account,
		}
		dao.GetAccountByAccount(oldAccount)
		if oldAccount.Id > 0 {
			service.JsonResponse(ctx, errors.ErrorCodeAccountAlreadyExist, nil)
			return
		}
		// 创建账号
		newAccount := &dao.Account{
			Account:    req.Account,
			Password:   service.GetPassword("123456"),
			IsSupper:   req.IsSupper,
			IsActivate: 1,
		}
		dao.CreateAccount(newAccount)
		service.JsonResponse(ctx, errors.CodeSuccess, nil)
	})
}

// Activate 账号激活
func (a *accountController) Activate(ctx *gin.Context) {
	req := &request.DetailBaseRequest{}
	if err := service.CheckValid(ctx, req); err != nil {
		service.JsonResponse(ctx, errors.ErrorCodeParamsError, nil)
		return
	}
	// 校验参数
	account := service.GetAccount(ctx)
	if req.Id == account.Id {
		service.JsonResponse(ctx, errors.ErrorCodeCanNotChangeSelf, nil)
		return
	}
	// 修改状态
	dao.AccountUpdate(req.Id, map[string]interface{}{
		"is_activate": 1,
	})
	service.JsonResponse(ctx, errors.CodeSuccess, nil)
}

// Deactivate 账号激活
func (a *accountController) Deactivate(ctx *gin.Context) {
	req := &request.DetailBaseRequest{}
	if err := service.CheckValid(ctx, req); err != nil {
		service.JsonResponse(ctx, errors.ErrorCodeParamsError, nil)
		return
	}
	// 校验参数
	account := service.GetAccount(ctx)
	if req.Id == account.Id {
		service.JsonResponse(ctx, errors.ErrorCodeCanNotChangeSelf, nil)
		return
	}
	// 修改状态
	dao.AccountUpdate(req.Id, map[string]interface{}{
		"is_activate": 0,
	})
	service.JsonResponse(ctx, errors.CodeSuccess, nil)
}

// ResetPassword 账号重置密码
func (a *accountController) ResetPassword(ctx *gin.Context) {

	req := &request.DetailBaseRequest{}
	if err := service.CheckValid(ctx, req); err != nil {
		service.JsonResponse(ctx, errors.ErrorCodeParamsError, nil)
		return
	}
	// 校验参数
	account := service.GetAccount(ctx)
	if req.Id == account.Id {
		service.JsonResponse(ctx, errors.ErrorCodeCanNotChangeSelf, nil)
		return
	}
	// 修改密码
	dao.AccountUpdate(req.Id, map[string]interface{}{
		"password": service.GetPassword("123456"),
	})
	service.JsonResponse(ctx, errors.CodeSuccess, nil)
}
