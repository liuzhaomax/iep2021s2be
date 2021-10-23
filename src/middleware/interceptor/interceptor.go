/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/7/24 17:31
 * @version     v1.0
 * @filename    interceptor.go
 * @description
 ***************************************************************************/
package interceptor

import "github.com/google/wire"

var InterceptorSet = wire.NewSet(
	AuthSet,
)
