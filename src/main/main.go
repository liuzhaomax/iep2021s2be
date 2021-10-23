/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/7/27 8:13
 * @version     v1.0
 * @filename    main.go
 * @description
 ***************************************************************************/
package main

import (
	"cms/src/app"
	"cms/src/core"
	"context"
)

func main() {
	const StaticDir = "/static"
	const ConfigDir = "/config"
	const ConfigFile = "config.json"
	ctx := context.Background()
	projectPath := core.GetProjectPath()
	app.Launch(ctx,
		app.SetStaticDir(projectPath+StaticDir),
		app.SetConfigDir(projectPath+ConfigDir),
		app.SetConfigFile(ConfigFile))
}
