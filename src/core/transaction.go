/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/8/8 11:38
 * @version     v1.0
 * @filename    transaction.go
 * @description
 ***************************************************************************/
package core

import "C"
import (
	"context"
	"github.com/jinzhu/gorm"
)

type transCtx struct{}

func NewTrans(ctx context.Context, tx interface{}) context.Context {
	return context.WithValue(ctx, transCtx{}, tx)
}

func GetTrans(ctx context.Context) (interface{}, bool) {
	v := ctx.Value(transCtx{})
	return v, v != nil
}

func ExecTrans(ctx context.Context, db *gorm.DB, fn func(context.Context) error) error {
	if _, ok := GetTrans(ctx); ok {
		return fn(ctx)
	}
	return db.Transaction(func(tx *gorm.DB) error {
		return fn(NewTrans(ctx, tx))
	})
}
