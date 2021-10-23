/****************************************************************************
 * @copyright   Team A28
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/9/5 23:17
 * @version     v1.0
 * @filename    h_stathub.go
 * @description
 ***************************************************************************/
package handler

import (
	"cms/src/core"
	"cms/src/module.epau/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"net/http"
)

var StatHubSet = wire.NewSet(wire.Struct(new(HStatHub), "*"))

type HStatHub struct {
	BStatHub *service.BStatHub
	IRes     core.IResponse
}

func (hStatHub *HStatHub) GetFEDataSource(ctx *gin.Context) {
	projectPath := core.GetProjectPath()
	filName := projectPath + "/static/ep_stat.csv"
	content, err := hStatHub.BStatHub.ReadData(filName)
	if err != nil {
		hStatHub.IRes.ResFail(ctx, http.StatusNotFound, err)
	}
	hStatHub.IRes.ResSuccess(ctx, content)
}

func (hStatHub *HStatHub) GetDataSource(ctx *gin.Context) {
	hStatHub.IRes.ResSuccess(ctx, "ok -> getDataSource")
}

func (hStatHub *HStatHub) PostDataSource(ctx *gin.Context) {
	fmt.Println(ctx.Request.Body)
	hStatHub.IRes.ResSuccess(ctx, "ok -> postDataSource")
}
