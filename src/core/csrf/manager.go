/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/7/10 2:27
 * @version     v1.0
 * @filename    manager.go
 * @description
 ***************************************************************************/
package csrf

import (
	"fmt"
	"sync"
	"time"
)

type Manager struct {
	lock        sync.Mutex
	provider    Provider
	maxLifeTime int64
}

var providers = make(map[string]Provider)

func NewSessionManager(providerName string, maxLifeTime int64) (*Manager, error) {
	provider, ok := providers[providerName]
	if !ok {
		return nil, fmt.Errorf("Session: unknown provide %q ", providerName)
	}
	return &Manager{provider: provider, maxLifeTime: maxLifeTime}, nil
}

func Register(name string, provider Provider) {
	if provider == nil {
		panic("Session: Register provider is nil")
	}
	if _, ok := providers[name]; ok {
		panic("Session: Register called twice for provider " + name)
	}
	providers[name] = provider
}

func (manager *Manager) SessionId(sid string) string {
	//return core.BASE64EncodeStr(sid)
	return sid
}

func (manager *Manager) SessionRead(sid string) (session Session) {
	session, _ = manager.provider.SessionRead(sid)
	return session
}

func (manager *Manager) SessionStart(sid string) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cipherSid := manager.SessionId(sid)
	session, _ = manager.provider.SessionInit(cipherSid)
	return session
}

func (manager *Manager) SessionDestroy(cipherSid string) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	//sid, _ := core.BASE64DecodeStr(cipherSid)
	_ = manager.provider.SessionDestroy(cipherSid)
}

func (manager *Manager) SessionGC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionGC(manager.maxLifeTime)
	time.AfterFunc(time.Duration(manager.maxLifeTime), func() { manager.SessionGC() })
}
