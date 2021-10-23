/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/8/1 3:07
 * @version     v1.0
 * @filename    web.go
 * @description
 ***************************************************************************/
package config

import (
	"cms/src/api"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
)

func init() {
	gin.DefaultWriter = colorable.NewColorableStdout()
	gin.ForceConsoleColor()
	//gin.DisableConsoleColor()
}

func InitGinEngine(iHandler api.IHandler) *gin.Engine {
	gin.SetMode(cfg.RunMode) // debug, test, release
	app := gin.Default()
	app.Use(LoggerToFile())
	iHandler.Register(app)
	return app
}
