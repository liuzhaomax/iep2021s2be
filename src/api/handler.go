/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/8/2 1:45
 * @version     v1.0
 * @filename    handler.go
 * @description
 ***************************************************************************/
package api

import (
	baseHandler "cms/src/module.base/handler"
	epauHandler "cms/src/module.epau/handler"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var APISet = wire.NewSet(wire.Struct(new(Handler), "*"), wire.Bind(new(IHandler), new(*Handler)))

type Handler struct {
	HandlerUser     *baseHandler.HUser
	HandlerPlantHub *epauHandler.HPlantHub
	HandlerStatHub  *epauHandler.HStatHub
	HandlerBlog     *epauHandler.HBlog
}

type IHandler interface {
	Register(app *gin.Engine)
}

func (handler *Handler) Register(app *gin.Engine) {
	handler.RegisterRouter(app)
}
