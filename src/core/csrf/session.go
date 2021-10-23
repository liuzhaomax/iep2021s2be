/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/7/29 10:47
 * @version     v1.0
 * @filename    session.go
 * @description
 ***************************************************************************/
package csrf

import (
	"container/list"
)

type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(ket interface{}) error
	SessionID() string
	SessionLastAccessedTime() int64
}

var sProvider = &SessionProvider{list: list.New()}

func init() {
	sProvider.sessions = make(map[string]*list.Element, 0)
	Register("sProvider", sProvider)
}

type SessionStore struct {
	sid              string
	LastAccessedTime int64
	value            map[interface{}]interface{} // session value
}

func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	err := sProvider.SessionUpdate(st.sid)
	if err != nil {
		return err
	}
	return nil
}

func (st *SessionStore) Get(key interface{}) interface{} {
	err := sProvider.SessionUpdate(st.sid)
	if err != nil {
		return err
	}
	if value, ok := st.value[key]; ok {
		return value
	} else {
		return nil
	}
}

func (st *SessionStore) Delete(key interface{}) error {
	delete(st.value, key)
	err := sProvider.SessionUpdate(st.sid)
	if err != nil {
		return err
	}
	return nil
}

func (st *SessionStore) SessionID() string {
	return st.sid
}

func (st *SessionStore) SessionLastAccessedTime() int64 {
	return st.LastAccessedTime
}
