/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/8/9 11:25
 * @version     v1.0
 * @filename    r_plantHub.go
 * @description
 ***************************************************************************/
package router

import (
	"cms/src/middleware/interceptor"
	"cms/src/module.base/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(handler *handler.HUser, group *gin.RouterGroup) {
	inter := interceptor.GetInstanceOfContext()

	routerHome := group.Group("/home")
	{
		routerHome.Use(inter.CheckTwoTokens())
		routerHome.Use(inter.CheckTokenWithinSession())

		routerHome.GET("", handler.GetHome)
		routerHome.POST("/register", handler.PostCreateUser)
	}
}
