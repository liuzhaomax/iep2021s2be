/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/9/16 0:22
 * @version     v1.0
 * @filename    r_blog.go
 * @description
 ***************************************************************************/
package router

import (
	"cms/src/middleware/interceptor"
	"cms/src/module.epau/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRouterBlog(handler *handler.HBlog, group *gin.RouterGroup) {
	cst := GetInstanceOfConstant()
	inter := interceptor.GetInstanceOfContext()

	routerFEBlog := group.Group(cst.Website1.Module4.Path)
	{
		routerFEBlog.GET(cst.Website1.Module4.Function1.Path+"/:limit", handler.GetFEBlogList)
		routerFEBlog.GET(cst.Website1.Module4.Function2.Path+"/:blogid", handler.GetFEBlog)
	}

	routerEpau := group.Group(cst.Website1.Path)
	{
		routerEpau.Use(inter.CheckTwoTokens())
		routerEpau.Use(inter.CheckTokenWithinSession())

		routerBlog := routerEpau.Group(cst.Website1.Module4.Path)
		{
			routerBlog.GET(cst.Website1.Module4.Function1.Path+"/read/:pagenum", handler.GetBlogList)
			routerBlog.GET(cst.Website1.Module4.Function1.Path+"/create", handler.GetCreateBlog)
			routerBlog.GET(cst.Website1.Module4.Function1.Path+"/update/:blogid", handler.GetBlog)
			routerBlog.POST(cst.Website1.Module4.Function1.Path+"/create", handler.PostBlog)
			routerBlog.PUT(cst.Website1.Module4.Function1.Path+"/update/:blogid", handler.PutBlog)
			routerBlog.DELETE(cst.Website1.Module4.Function1.Path+"/delete/:blogid", handler.DeleteBlog)
		}
	}
}
