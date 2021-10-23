// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package app

import (
	"cms/src/api"
	"cms/src/config"
	"cms/src/core"
	"cms/src/middleware/interceptor"
	"cms/src/module.base/handler"
	"cms/src/module.base/model"
	"cms/src/module.base/service"
	handler2 "cms/src/module.epau/handler"
	model2 "cms/src/module.epau/model"
	service2 "cms/src/module.epau/service"
)

// Injectors from wire.go:

func InitInjector() (*Injector, func(), error) {
	db, cleanup, err := InitDB()
	if err != nil {
		return nil, nil, err
	}
	user := &model.User{
		Tx: db,
	}
	bUser := &service.BUser{
		MUser: user,
	}
	response := &core.Response{}
	hUser := &handler.HUser{
		BUser: bUser,
		IRes:  response,
	}
	bPlantHub := &service2.BPlantHub{}
	hPlantHub := &handler2.HPlantHub{
		BPlantHub: bPlantHub,
		IRes:      response,
	}
	bStatHub := &service2.BStatHub{}
	hStatHub := &handler2.HStatHub{
		BStatHub: bStatHub,
		IRes:     response,
	}
	blog := &model2.Blog{
		Tx: db,
	}
	bBlog := &service2.BBlog{
		MBlog: blog,
	}
	hBlog := &handler2.HBlog{
		BBlog: bBlog,
		IRes:  response,
	}
	apiHandler := &api.Handler{
		HandlerUser:     hUser,
		HandlerPlantHub: hPlantHub,
		HandlerStatHub:  hStatHub,
		HandlerBlog:     hBlog,
	}
	engine := config.InitGinEngine(apiHandler)
	interceptorInterceptor := &interceptor.Interceptor{}
	injector := &Injector{
		Engine:      engine,
		DB:          db,
		Handler:     apiHandler,
		Interceptor: interceptorInterceptor,
	}
	return injector, func() {
		cleanup()
	}, nil
}