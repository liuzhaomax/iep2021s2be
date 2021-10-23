/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/8/5 7:19
 * @version     v1.0
 * @filename    b_user.go
 * @description
 ***************************************************************************/
package service

import (
	"cms/src/core"
	"cms/src/module.base/entity"
	"cms/src/module.base/model"
	"cms/src/module.base/schema"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"math"
	"strconv"
	"strings"
	"time"
)

var UserSet = wire.NewSet(wire.Struct(new(BUser), "*"))

type BUser struct {
	MUser *model.User
}

func (bUser *BUser) SetLoginCookie(ctx *gin.Context, loginInfo *schema.RegisterInfo) *core.Error {
	duration := time.Hour * 24 * 7 // a week
	userEmail, err := core.RSADecrypt(core.GetPrivateKey(), loginInfo.UserEmail)
	token, err := core.GenerateToken(userEmail, duration)
	if err != nil {
		return core.NewError(106, err)
	}
	cipherToken, err := core.RSAEncrypt(core.GetPublicKey(), token)
	if err != nil {
		return core.NewError(107, err)
	}
	targetDomain := core.GetFEAdminDomain()
	secure := core.GetParamsByProtocol()
	durationInt := int(duration) / int(math.Pow10(9))
	ctx.SetCookie(
		"TOKEN",
		cipherToken,
		durationInt,
		"/",
		targetDomain,
		secure,
		true)
	return nil
}

func (bUser *BUser) SetLoginJWT(ctx *gin.Context, loginInfo *schema.RegisterInfo) (string, string, *core.Error) {
	duration := time.Hour * 24 * 7 // a week
	userEmail, err := core.RSADecrypt(core.GetPrivateKey(), loginInfo.UserEmail)
	if err != nil {
		return "", "", core.NewError(108, err)
	}
	session := core.GetInstanceOfContext().SessionManager.SessionStart(userEmail)
	value := session.SessionID() + "|" + strconv.FormatInt(session.SessionLastAccessedTime(), 10)
	token, err := core.GenerateToken(value, duration)
	if err != nil {
		return "", "", core.NewError(106, err)
	}
	cipherToken, _ := core.RSAEncrypt(core.GetPublicKey(), token)
	return cipherToken, userEmail, nil
}

func (bUser *BUser) CheckLoginUser(ctx context.Context, registerInfo *schema.RegisterInfo) (bool, *core.Error) {
	var user = entity.User{}
	user.UserEmail, _ = core.RSADecrypt(core.GetPrivateKey(), registerInfo.UserEmail)
	user.UserPassword, _ = core.RSADecrypt(core.GetPrivateKey(), registerInfo.UserPassword)
	userQueried, err := bUser.MUser.QueryOneUserWithPassword(&user)
	if err != nil {
		return false, core.NewError(900, nil)
	}
	if userQueried == nil {
		return false, core.NewError(104, nil)
	}
	if userQueried.UserPassword != user.UserPassword {
		return false, core.NewError(105, nil)
	}
	return true, nil
}

func (bUser *BUser) ClearLoginCookie(ctx *gin.Context) (string, *core.Error) {
	cookieToken, _ := ctx.Cookie("TOKEN")
	cookieToken, _ = core.RSADecrypt(core.GetPrivateKey(), cookieToken)
	cookieTokenEmail, _ := core.ParseToken(cookieToken)
	targetDomain := core.GetFEAdminDomain()
	secure := core.GetParamsByProtocol()
	ctx.SetCookie(
		"TOKEN",
		"",
		1,
		"/",
		targetDomain,
		secure,
		true)
	return cookieTokenEmail, nil
}

func (bUser *BUser) ClearLoginSession(ctx *gin.Context, sessionId string) {
	manager := core.GetInstanceOfContext().SessionManager
	manager.SessionDestroy(sessionId)
}

func (bUser *BUser) CreateOneUser(ctx context.Context, registerInfo *schema.RegisterInfo) *core.Error {
	var user = entity.User{}
	user.UserEmail = registerInfo.UserEmail
	user.UserPassword = registerInfo.UserPassword
	user.UserNickName = bUser.generateDefaultNickName(registerInfo.UserEmail)
	user.UserCreateTime = time.Now()
	_, err := bUser.MUser.QueryOneUser(&user)
	if err == nil {
		return core.NewError(100, nil)
	}
	err = core.ExecTrans(ctx, bUser.MUser.Tx, func(ctx context.Context) error {
		err := bUser.MUser.CreateOneUser(&user)
		if err != nil {
			return core.NewError(100, err)
		}
		return nil
	})
	if err != nil {
		return core.NewError(900, err)
	}
	_, err = bUser.MUser.QueryOneUser(&user)
	if err != nil {
		return core.NewError(900, err)
	}
	return nil
}

func (bUser *BUser) generateDefaultNickName(userEmail string) string {
	emailName := strings.Split(userEmail, "@")[0]
	var emailNameFirstPart string
	if len(emailName) >= 5 {
		emailNameFirstPart = emailName[0:5]
	} else {
		emailNameFirstPart = emailName
	}
	lastUserId := core.GetInstanceOfContext().LastUserId
	return emailNameFirstPart + "_" + strconv.Itoa(int(lastUserId)+1)
}
