/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/8/2 1:10
 * @version     v1.0
 * @filename    r_plantHub.go
 * @description
 ***************************************************************************/
package api

import (
	mw "cms/src/middleware"
	base "cms/src/module.base/router"
	epau "cms/src/module.epau/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler *Handler) RegisterRouter(app *gin.Engine) {
	app.NoRoute(handler.HandlerUser.GetNoRoute)
	app.Use(mw.Cors())

	router := app.Group("")
	{
		router.StaticFS("/static", http.Dir("./static"))
		router.GET("/", handler.HandlerUser.GetIndex)
		router.GET("/login", handler.HandlerUser.GetPuk)
		router.POST("/login", handler.HandlerUser.PostLogin)
		router.DELETE("/logout", handler.HandlerUser.DeleteLogout)
		base.RegisterRouter(handler.HandlerUser, router)
		epau.RegisterRouterPlantHub(handler.HandlerPlantHub, router)
		epau.RegisterRouterStatHub(handler.HandlerStatHub, router)
		epau.RegisterRouterBlog(handler.HandlerBlog, router)
	}
}
