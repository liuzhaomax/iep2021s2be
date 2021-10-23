/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/8/8 20:54
 * @version     v1.0
 * @filename    entity.go
 * @description
 ***************************************************************************/
package entity

import "github.com/jinzhu/copier"

func Transform(ent, sch interface{}) error {
	return copier.Copy(sch, ent)
}
