/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/7/31 3:10
 * @version     v1.0
 * @filename    r_plantHub.go
 * @description
 ***************************************************************************/
package router

import (
	"cms/src/middleware/interceptor"
	"cms/src/module.epau/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRouterPlantHub(handler *handler.HPlantHub, group *gin.RouterGroup) {
	cst := GetInstanceOfConstant()
	inter := interceptor.GetInstanceOfContext()

	routerFEPlantHub := group.Group(cst.Website1.Module1.Path)
	{
		routerFEPlantHub.GET(cst.Website1.Module1.Function1.Path, handler.GetFEDataSource)
	}

	routerEpau := group.Group(cst.Website1.Path)
	{
		routerEpau.Use(inter.CheckTwoTokens())
		routerEpau.Use(inter.CheckTokenWithinSession())

		routerPlantHub := routerEpau.Group(cst.Website1.Module1.Path)
		{
			routerPlantHub.GET(cst.Website1.Module1.Function1.Path, handler.GetDataSource)
			routerPlantHub.POST(cst.Website1.Module1.Function1.Path, handler.PostDataSource)
		}
	}
}
