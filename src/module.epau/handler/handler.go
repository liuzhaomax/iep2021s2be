/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/8/6 3:29
 * @version     v1.0
 * @filename    handler.go
 * @description
 ***************************************************************************/
package handler

import (
	"cms/src/core"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var HandlerSet = wire.NewSet(
	PlantHubSet,
	StatHubSet,
	BlogSet,
)

func GetUserEmail(ctx *gin.Context) string {
	cookieToken, _ := ctx.Cookie("TOKEN")
	cookieToken, _ = core.RSADecrypt(core.GetPrivateKey(), cookieToken)
	cookieTokenEmail, _ := core.ParseToken(cookieToken)
	return cookieTokenEmail
}
