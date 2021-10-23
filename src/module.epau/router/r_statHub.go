/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/9/5 23:33
 * @version     v1.0
 * @filename    r_statHub.go
 * @description
 ***************************************************************************/
package router

import (
	"cms/src/middleware/interceptor"
	"cms/src/module.epau/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRouterStatHub(handler *handler.HStatHub, group *gin.RouterGroup) {
	cst := GetInstanceOfConstant()
	inter := interceptor.GetInstanceOfContext()

	routerFEStatHub := group.Group(cst.Website1.Module2.Path)
	{
		routerFEStatHub.GET(cst.Website1.Module2.Function1.Path, handler.GetFEDataSource)
	}

	routerEpau := group.Group(cst.Website1.Path)
	{
		routerEpau.Use(inter.CheckTwoTokens())
		routerEpau.Use(inter.CheckTokenWithinSession())

		routerStatHub := routerEpau.Group(cst.Website1.Module2.Path)
		{
			routerStatHub.GET(cst.Website1.Module2.Function1.Path, handler.GetDataSource)
			routerStatHub.POST(cst.Website1.Module2.Function1.Path, handler.PostDataSource)
		}
	}
}
