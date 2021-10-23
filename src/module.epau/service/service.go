/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/8/6 3:33
 * @version     v1.0
 * @filename    service.go
 * @description
 ***************************************************************************/
package service

import (
	"github.com/google/wire"
)

var ServiceSet = wire.NewSet(
	PlantHubSet,
	StatHubSet,
	BlogSet,
)
