/****************************************************************************
 * @copyright   Team A28
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/7/31 2:51
 * @version     v1.0
 * @filename    h_plantHub.go
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

var PlantHubSet = wire.NewSet(wire.Struct(new(HPlantHub), "*"))

type HPlantHub struct {
	BPlantHub *service.BPlantHub
	IRes      core.IResponse
}

func (hPlantHub *HPlantHub) GetFEDataSource(ctx *gin.Context) {
	projectPath := core.GetProjectPath()
	filName := projectPath + "/static/endangeredplants.csv"
	content, err := hPlantHub.BPlantHub.ReadData(filName)
	if err != nil {
		hPlantHub.IRes.ResFail(ctx, http.StatusNotFound, err)
	}
	hPlantHub.IRes.ResSuccess(ctx, content)
}

func (hPlantHub *HPlantHub) GetDataSource(ctx *gin.Context) {
	hPlantHub.IRes.ResSuccess(ctx, "ok -> getDataSource")
}

func (hPlantHub *HPlantHub) PostDataSource(ctx *gin.Context) {
	fmt.Println(ctx.Request.Body)
	hPlantHub.IRes.ResSuccess(ctx, "ok -> postDataSource")
}
