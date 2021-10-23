/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/8/6 3:29
 * @version     v1.0
 * @filename    handler.go
 * @description
 ***************************************************************************/
package handler

import (
	"github.com/google/wire"
)

var HandlerSet = wire.NewSet(
	UserSet,
)
