/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/8/6 3:35
 * @version     v1.0
 * @filename    model.go
 * @description
 ***************************************************************************/
package model

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

var ModelSet = wire.NewSet(
	BlogSet,
)

func addEL(tx *gorm.DB) *gorm.DB {
	return tx.Set("gorm:query_option", "FOR UPDATE")
}

//func addELWhenDelete(tx *gorm.DB) *gorm.DB {
//	return tx.Set("gorm:delete_option", "OPTION (OPTIMIZE FOR UNKNOWN)")
//}

//func GetDBWhenQuery(ctx context.Context, db *gorm.DB) *gorm.DB {
//	trans, ok := core.GetTrans(ctx)
//	tx, ok1 := trans.(*gorm.DB)
//	if ok && ok1 {
//		return tx
//	}
//	return db
//}
//
//func GetDBWhenQueryWithModel(ctx context.Context, db *gorm.DB, entity interface{}) *gorm.DB {
//	return GetDBWhenQuery(ctx, db).Model(entity)
//}
//
//func GetDBWhenEL(ctx context.Context, db *gorm.DB) *gorm.DB {
//	trans, ok := core.GetTrans(ctx)
//	tx, ok1 := trans.(*gorm.DB)
//	if ok && ok1 {
//		cfg := config.GetInstanceOfConfig()
//		if cfg.DB.Type == "mysql" || cfg.DB.Type == "postgresql" {
//			tx = tx.Set("gorm:query_option", "FOR UPDATE")
//		}
//		return tx
//	}
//	return db
//}
//
//func GetDBWhenELWithModel(ctx context.Context, db *gorm.DB, entity interface{}) *gorm.DB {
//	return GetDBWhenEL(ctx, db).Model(entity)
//}
