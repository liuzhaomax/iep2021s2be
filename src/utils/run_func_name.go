/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/7/31 7:44
 * @version     v1.0
 * @filename    run_func_name.go
 * @description
 ***************************************************************************/
package utils

import "runtime"

func RunFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	function := runtime.FuncForPC(pc[0])
	return " - " + function.Name()
}
