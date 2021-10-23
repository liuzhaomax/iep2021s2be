/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/8/3 1:56
 * @version     v1.0
 * @filename    i_auth.go
 * @description
 ***************************************************************************/
package interceptor

import (
	"cms/src/core"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var AuthSet = wire.NewSet(wire.Struct(new(Interceptor), "*"))

var interceptor *Interceptor
var once sync.Once

func init() {
	once.Do(func() {
		interceptor = &Interceptor{}
	})
}

func GetInstanceOfContext() *Interceptor {
	return interceptor
}

type Interceptor struct{}

// emails in tokens need to be equal
func (inter *Interceptor) CheckTwoTokens() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// token in req header
		headerToken := ctx.Request.Header.Get("Authorization")
		headerToken, _ = core.RSADecrypt(core.GetPrivateKey(), headerToken)
		headerTokenEmail, _ := core.ParseToken(headerToken)
		headerTokenEmail = strings.Split(headerTokenEmail, "|")[0]
		// token in req cookie
		cookieToken, _ := ctx.Cookie("TOKEN")
		cookieToken, _ = core.RSADecrypt(core.GetPrivateKey(), cookieToken)
		cookieTokenEmail, _ := core.ParseToken(cookieToken)
		// checking tokens info
		if headerTokenEmail != cookieTokenEmail || headerTokenEmail == "" || cookieTokenEmail == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, core.NewError(110, nil))
		}
		ctx.Next()
	}
}

// emails plus durations in token and session need to be equal
func (inter *Interceptor) CheckTokenWithinSession() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// token in req header
		headerToken := ctx.Request.Header.Get("Authorization")
		headerToken, _ = core.RSADecrypt(core.GetPrivateKey(), headerToken)
		headerTokenEmailPlusDue, _ := core.ParseToken(headerToken)
		headerTokenEmail := strings.Split(headerTokenEmailPlusDue, "|")[0]
		var headerTokenDueStr string
		var headerTokenDue int64
		if len(strings.Split(headerTokenEmailPlusDue, "|")) > 1 {
			headerTokenDueStr = strings.Split(headerTokenEmailPlusDue, "|")[1]
			headerTokenDue, _ = strconv.ParseInt(headerTokenDueStr, 10, 64)
		}
		// token in session
		manager := core.GetInstanceOfContext().SessionManager
		session := manager.SessionRead(headerTokenEmail)
		sessionEmail := session.SessionID()
		sessionDue := session.SessionLastAccessedTime()
		// check tokens info
		if headerTokenEmail != sessionEmail || headerTokenEmail == "" || headerTokenDue != sessionDue {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, core.NewError(110, nil))
		}
		ctx.Next()
	}
}
