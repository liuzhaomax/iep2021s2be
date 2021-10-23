/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/7/30 3:52
 * @version     v1.0
 * @filename    cors.go
 * @description
 ***************************************************************************/
package middleware

import (
	"cms/src/core"
	"cms/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	var corsWhiteList = []string{
		core.GetFEAdminProtocol() + "://" + core.GetFEAdminHost() + ":" + core.GetFEAdminPort(),
		core.GetFEAdminProtocol() + "://" + core.GetFEAdminDomain(),
		core.GetFEMainProtocol() + "://" + core.GetFEMainHost() + ":" + core.GetFEMainPort(),
		core.GetFEMainProtocol() + "://" + core.GetFEMainDomain(),
		core.GetFEMainProtocol() + "://www." + core.GetFEMainDomain(),
		core.GetFEMainProtocol() + "://iter1." + core.GetFEMainDomain(),
		core.GetFEMainProtocol() + "://iter2." + core.GetFEMainDomain(),
		core.GetFEMainProtocol() + "://iter3." + core.GetFEMainDomain(),
	}
	return func(ctx *gin.Context) {
		if utils.In(corsWhiteList, ctx.Request.Header.Get("Origin")) {
			ctx.Header("Access-Control-Allow-Origin", ctx.Request.Header.Get("Origin"))
		}
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token, Set-Cookie, X-Requested-With, Access-Control-Allow-Origin, Content-Security-Policy")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		// let go all options request
		method := ctx.Request.Method
		if method == "OPTIONS" {
			ctx.Header("Access-Control-Max-Age", "86400") // one day
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		ctx.Next()
	}
}
