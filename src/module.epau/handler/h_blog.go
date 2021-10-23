/****************************************************************************
 * @copyright   Team A28
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/9/16 0:24
 * @version     v1.0
 * @filename    h_blog.go
 * @description
 ***************************************************************************/
package handler

import (
	"cms/src/core"
	"cms/src/module.epau/schema"
	"cms/src/module.epau/service"
	"cms/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	logger "github.com/sirupsen/logrus"
	"net/http"
)

var BlogSet = wire.NewSet(wire.Struct(new(HBlog), "*"))

type HBlog struct {
	BBlog *service.BBlog
	IRes  core.IResponse
}

func (hBlog *HBlog) GetFEBlogList(ctx *gin.Context) {
	blogList, _ := hBlog.BBlog.ReadFEBlogList(ctx)
	hBlog.IRes.ResSuccess(ctx, blogList)
}

func (hBlog *HBlog) GetFEBlog(ctx *gin.Context) {
	blog, _ := hBlog.BBlog.ReadBlog(ctx)
	err := hBlog.BBlog.IncreaseViews(ctx)
	if err != nil {
		hBlog.IRes.ResFail(ctx, http.StatusBadRequest, core.NewError(301, err))
	}
	hBlog.IRes.ResSuccess(ctx, blog)
}

func (hBlog *HBlog) GetBlogList(ctx *gin.Context) {
	blogList, _ := hBlog.BBlog.ReadBlogList(ctx)
	logger.Info(GetUserEmail(ctx), utils.RunFuncName())
	hBlog.IRes.ResSuccess(ctx, blogList)
}

func (hBlog *HBlog) GetCreateBlog(ctx *gin.Context) {
	logger.Info(GetUserEmail(ctx), utils.RunFuncName())
	hBlog.IRes.ResSuccess(ctx, "ok")
}

func (hBlog *HBlog) GetBlog(ctx *gin.Context) {
	blog, _ := hBlog.BBlog.ReadBlog(ctx)
	logger.Info(GetUserEmail(ctx), utils.RunFuncName())
	hBlog.IRes.ResSuccess(ctx, blog)
}

func (hBlog *HBlog) PostBlog(ctx *gin.Context) {
	var blogWithContent schema.BlogWithContent
	if err := ctx.ShouldBind(&blogWithContent); err != nil {
		hBlog.IRes.ResFail(ctx, http.StatusBadRequest, core.NewError(901, err))
	} else {
		err = hBlog.BBlog.CreateBlog(ctx, &blogWithContent)
		if err != nil {
			hBlog.IRes.ResFail(ctx, http.StatusBadRequest, core.NewError(300, err))
		}
		logger.Info(GetUserEmail(ctx), utils.RunFuncName())
		hBlog.IRes.ResSuccess(ctx, "ok")
	}
}

func (hBlog *HBlog) PutBlog(ctx *gin.Context) {
	var blogWithContent schema.BlogWithContent
	if err := ctx.ShouldBind(&blogWithContent); err != nil {
		hBlog.IRes.ResFail(ctx, http.StatusBadRequest, core.NewError(901, err))
	} else {
		err = hBlog.BBlog.UpdateBlog(ctx, &blogWithContent)
		if err != nil {
			hBlog.IRes.ResFail(ctx, http.StatusBadRequest, core.NewError(301, err))
		}
		logger.Info(GetUserEmail(ctx), utils.RunFuncName())
		hBlog.IRes.ResSuccess(ctx, "ok")
	}
}

func (hBlog *HBlog) DeleteBlog(ctx *gin.Context) {
	err := hBlog.BBlog.DeleteBlog(ctx)
	if err != nil {
		hBlog.IRes.ResFail(ctx, http.StatusBadRequest, core.NewError(302, err))
	}
	logger.Info(GetUserEmail(ctx), utils.RunFuncName())
	hBlog.IRes.ResSuccess(ctx, "ok")
}
