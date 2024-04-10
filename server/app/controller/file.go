package controller

import (
	"config_tools/app/errors"
	"config_tools/app/service"
	"config_tools/tools/net/middleware"
	"fmt"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kennycch/gotools/general"
)

type fileController struct{}

var (
	fileControllers = fileController{}
)

func FileRoute(route *gin.Engine) {
	route.Static("/file", "./file")
	// 需要校验登录的
	lRoute := route.Group("", middleware.CheckLogin())
	{
		// 上传图片
		lRoute.POST("files/uploadImg", fileControllers.UploadImg)
	}
}

// UploadImg 上传图片
func (f *fileController) UploadImg(ctx *gin.Context) {
	// 读取图片
	file, err := ctx.FormFile("img")
	if err != nil {
		service.JsonResponse(ctx, errors.ErrorCodeParamsError, nil)
		return
	}
	// 校验文件后缀
	ext := path.Ext(file.Filename)
	if !general.InArray([]string{".jpg", ".png", ".jpeg"}, ext) {
		service.JsonResponse(ctx, errors.ErrorCodeFileExtNotMacth, nil)
		return
	}
	// 生成图片路径
	imgPath := fmt.Sprintf("./file/img/%s/%s%s",
		time.Now().Format("20060102"),
		general.GetUniqueId(),
		ext,
	)
	if err = ctx.SaveUploadedFile(file, imgPath); err != nil {
		service.JsonResponse(ctx, errors.CodeServerError, nil)
		return
	}
	service.JsonResponse(ctx, errors.CodeSuccess, gin.H{
		"path": string(imgPath[1:]),
	})
}
