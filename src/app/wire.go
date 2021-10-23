// +build wireinject

/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/8/6 0:19
 * @version     v1.0
 * @filename    wire.go
 * @description
 ***************************************************************************/
package app

import (
	"cms/src/api"
	"cms/src/config"
	"cms/src/core"
	"cms/src/middleware/interceptor"
	baseHandler "cms/src/module.base/handler"
	baseModel "cms/src/module.base/model"
	baseService "cms/src/module.base/service"
	epauHandler "cms/src/module.epau/handler"
	epauModel "cms/src/module.epau/model"
	epauService "cms/src/module.epau/service"
	"github.com/google/wire"
)

func InitInjector() (*Injector, func(), error) {
	wire.Build(
		InitDB,
		config.InitGinEngine,
		api.APISet,
		core.ResponseSet,
		interceptor.InterceptorSet,
		baseHandler.HandlerSet,
		baseService.ServiceSet,
		baseModel.ModelSet,
		epauHandler.HandlerSet,
		epauService.ServiceSet,
		epauModel.ModelSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
