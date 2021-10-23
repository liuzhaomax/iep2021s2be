/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/7/10 2:32
 * @version     v1.0
 * @filename    provider.go
 * @description
 ***************************************************************************/
package csrf

import (
	"container/list"
	"sync"
	"time"
)

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

type SessionProvider struct {
	lock     sync.Mutex
	sessions map[string]*list.Element
	list     *list.List // gc
}

func (sp *SessionProvider) SessionInit(sid string) (Session, error) {
	sp.lock.Lock()
	defer sp.lock.Unlock()
	value := make(map[interface{}]interface{}, 0)
	newSession := &SessionStore{sid: sid, LastAccessedTime: time.Now().Unix(), value: value}
	element := sp.list.PushBack(newSession)
	sp.sessions[sid] = element
	return newSession, nil
}

func (sp *SessionProvider) SessionRead(sid string) (Session, error) {
	if element, ok := sp.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		sess, err := sp.SessionInit(sid)
		return sess, err
	}
}

func (sp *SessionProvider) SessionDestroy(sid string) error {
	if element, ok := sp.sessions[sid]; ok {
		delete(sp.sessions, sid)
		sp.list.Remove(element)
		return nil
	}
	return nil
}

func (sp *SessionProvider) SessionUpdate(sid string) error {
	sp.lock.Lock()
	defer sp.lock.Unlock()
	if element, ok := sp.sessions[sid]; ok {
		element.Value.(*SessionStore).LastAccessedTime = time.Now().Unix()
		sp.list.MoveToFront(element)
		return nil
	}
	return nil
}

func (sp *SessionProvider) SessionGC(maxLifeTime int64) {
	sp.lock.Lock()
	defer sp.lock.Unlock()
	for {
		element := sp.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*SessionStore).LastAccessedTime + maxLifeTime) <
			time.Now().Unix() {
			sp.list.Remove(element)
			delete(sp.sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}
