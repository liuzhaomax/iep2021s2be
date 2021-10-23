/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/8/5 4:24
 * @version     v1.0
 * @filename    response.go
 * @description
 ***************************************************************************/
package core

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"net/http"
)

var ResponseSet = wire.NewSet(wire.Struct(new(Response), "*"), wire.Bind(new(IResponse), new(*Response)))

type Response struct{}

type IResponse interface {
	ResSuccess(ctx *gin.Context, sth interface{})
	ResFail(ctx *gin.Context, code int, err *Error)
}

func (res *Response) ResSuccess(ctx *gin.Context, sth interface{}) {
	res.ResJson(ctx, http.StatusOK, sth)
}

func (res *Response) ResJson(ctx *gin.Context, status int, sth interface{}) {
	ctx.JSON(status, sth)
	ctx.Abort()
}

func (res *Response) ResFail(ctx *gin.Context, code int, err *Error) {
	res.ResError(ctx, code, err)
}

func (res *Response) ResError(ctx *gin.Context, status int, err *Error) {
	res.ResJson(ctx, status, err)
}
