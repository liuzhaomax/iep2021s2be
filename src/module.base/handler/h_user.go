/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/8/2 2:06
 * @version     v1.0
 * @filename    h_user.go
 * @description
 ***************************************************************************/
package handler

import (
	"cms/src/core"
	"cms/src/module.base/schema"
	"cms/src/module.base/service"
	"cms/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	logger "github.com/sirupsen/logrus"
	"net/http"
)

var UserSet = wire.NewSet(wire.Struct(new(HUser), "*"))

type HUser struct {
	BUser *service.BUser
	IRes  core.IResponse
}

func (hUser *HUser) GetNoRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{"res": "404"})
}

func (hUser *HUser) GetIndex(ctx *gin.Context) {
	hUser.IRes.ResSuccess(ctx, "Hello EP Fans")
}

func (hUser *HUser) GetPuk(ctx *gin.Context) {
	puk := core.GetPublicKeyStr()
	hUser.IRes.ResSuccess(ctx, gin.H{"puk": puk})
}

func (hUser *HUser) PostLogin(ctx *gin.Context) {
	var loginInfo schema.RegisterInfo
	if err := ctx.ShouldBind(&loginInfo); err != nil {
		hUser.IRes.ResFail(ctx, http.StatusBadRequest, core.NewError(901, nil))
	} else {
		exist, err := hUser.BUser.CheckLoginUser(ctx.Request.Context(), &loginInfo)
		if exist == false {
			hUser.IRes.ResFail(ctx, http.StatusUnauthorized, err)
		} else {
			err := hUser.BUser.SetLoginCookie(ctx, &loginInfo)
			cipherToken, userEmail, err := hUser.BUser.SetLoginJWT(ctx, &loginInfo)
			if err != nil {
				hUser.IRes.ResFail(ctx, http.StatusInternalServerError, err)
			} else {
				logger.Info(userEmail, utils.RunFuncName())
				hUser.IRes.ResSuccess(ctx, cipherToken)
			}
		}
	}
}

func (hUser *HUser) DeleteLogout(ctx *gin.Context) {
	userEmail, _ := hUser.BUser.ClearLoginCookie(ctx)
	hUser.BUser.ClearLoginSession(ctx, userEmail)
	logger.Info(userEmail, utils.RunFuncName())
	hUser.IRes.ResSuccess(ctx, "ok")
}

func (hUser *HUser) GetHome(ctx *gin.Context) {
	hUser.IRes.ResSuccess(ctx, "ok -> getHome")
}

func (hUser *HUser) PostCreateUser(ctx *gin.Context) {
	var registerInfo schema.RegisterInfo
	if err := ctx.ShouldBind(&registerInfo); err != nil {
		hUser.IRes.ResFail(ctx, http.StatusBadRequest, core.NewError(901, nil))
	} else {
		err := hUser.BUser.CreateOneUser(ctx.Request.Context(), &registerInfo)
		if err != nil {
			hUser.IRes.ResFail(ctx, http.StatusInternalServerError, err)
		} else {
			hUser.IRes.ResSuccess(ctx, "ok")
		}
	}
}
