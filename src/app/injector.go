/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/8/5 10:09
 * @version     v1.0
 * @filename    injector.go
 * @description
 ***************************************************************************/
package app

import (
	"cms/src/api"
	"cms/src/middleware/interceptor"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	Engine      *gin.Engine
	DB          *gorm.DB
	Handler     *api.Handler
	Interceptor *interceptor.Interceptor
}
