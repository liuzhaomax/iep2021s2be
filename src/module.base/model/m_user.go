/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/8/3 6:06
 * @version     v1.0
 * @filename    m_blog.go
 * @description
 ***************************************************************************/
package model

import (
	"cms/src/core"
	"cms/src/module.base/entity"
	"cms/src/module.base/schema"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

var UserSet = wire.NewSet(wire.Struct(new(User), "*"))

type User struct {
	Tx *gorm.DB
}

func (mUser *User) QueryOneUser(user *entity.User) (*schema.User, error) {
	queriedUser := &entity.User{}
	mUser.Tx.Where("user_email=?", user.UserEmail).First(queriedUser)
	if int(queriedUser.UserId) == 0 {
		return nil, core.NewError(104, nil)
	}
	sUser := queriedUser.EntityToSchemaUser()
	return sUser, nil
}

func (mUser *User) QueryOneUserWithPassword(user *entity.User) (*schema.UserWithPassword, error) {
	queriedUser := &entity.User{}
	mUser.Tx.Where("user_email=?", user.UserEmail).First(queriedUser)
	if int(queriedUser.UserId) == 0 {
		return nil, core.NewError(104, nil)
	}
	sUser := queriedUser.EntityToSchemaUserWithPassword()
	return sUser, nil
}

func (mUser *User) CreateOneUser(user *entity.User) error {
	result := addEL(mUser.Tx).Create(&user)
	return result.Error
}
