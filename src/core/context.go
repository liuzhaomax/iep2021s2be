/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/7/30 23:26
 * @version     v1.0
 * @filename    context.go
 * @description Context singleton
 ***************************************************************************/
package core

import (
	"cms/src/core/csrf"
	"crypto/rsa"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

var ctx *Context
var once sync.Once

func init() {
	once.Do(func() {
		ctx = &Context{}
	})
}

func GetInstanceOfContext() *Context {
	return ctx
}

type Context struct {
	mux            sync.Mutex
	RunMode        string
	Http           Http
	FEAdmin        FEAdmin
	FEMain         FEMain
	JWTSecret      string
	PrivateKey     *rsa.PrivateKey
	PublicKey      *rsa.PublicKey
	PublicKeyStr   string
	LastUserId     uint
	SessionManager *csrf.Manager
}

type Http struct {
	Protocol string
	Host     string
}

type FEAdmin struct {
	LocalProtocol  string
	RemoteProtocol string
	Domain         string
	Host           string
	Port           string
}

type FEMain struct {
	LocalProtocol  string
	RemoteProtocol string
	Domain         string
	Host           string
	Port           string
}

func GetProjectPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	indexWithoutFileName := strings.LastIndex(path, string(os.PathSeparator))
	indexWithoutLastPath := strings.LastIndex(path[:indexWithoutFileName], string(os.PathSeparator))
	return strings.Replace(path[:indexWithoutLastPath], "\\", "/", -1)
}

func GetPublicKey() *rsa.PublicKey {
	return GetInstanceOfContext().PublicKey
}

func GetPublicKeyStr() string {
	return GetInstanceOfContext().PublicKeyStr
}

func GetPrivateKey() *rsa.PrivateKey {
	return GetInstanceOfContext().PrivateKey
}

func GetFEAdminProtocol() string {
	if GetInstanceOfContext().RunMode == "release" {
		return GetInstanceOfContext().FEAdmin.RemoteProtocol
	} else {
		return GetInstanceOfContext().FEAdmin.LocalProtocol
	}
}

func GetFEAdminDomain() string {
	if GetInstanceOfContext().RunMode == "release" {
		return GetInstanceOfContext().FEAdmin.Domain
	} else {
		return GetInstanceOfContext().FEAdmin.Host
	}
}

func GetFEAdminHost() string {
	return GetInstanceOfContext().FEAdmin.Host
}

func GetFEAdminPort() string {
	return GetInstanceOfContext().FEAdmin.Port
}

func GetFEMainProtocol() string {
	if GetInstanceOfContext().RunMode == "release" {
		return GetInstanceOfContext().FEMain.RemoteProtocol
	} else {
		return GetInstanceOfContext().FEMain.LocalProtocol
	}
}

func GetFEMainDomain() string {
	if GetInstanceOfContext().RunMode == "release" {
		return GetInstanceOfContext().FEMain.Domain
	} else {
		return GetInstanceOfContext().FEMain.Host
	}
}

func GetFEMainHost() string {
	return GetInstanceOfContext().FEMain.Host
}

func GetFEMainPort() string {
	return GetInstanceOfContext().FEMain.Port
}

func GetParamsByProtocol() bool {
	protocol := GetFEAdminProtocol()
	var secure bool
	switch protocol {
	case "http":
		secure = false
	case "https":
		secure = true
	}
	return secure
}
