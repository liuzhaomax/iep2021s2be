/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/8/3 4:05
 * @version     v1.0
 * @filename    e_blog.go
 * @description
 ***************************************************************************/
package entity

import (
	"cms/src/module.base/schema"
	"time"
)

type User struct {
	UserId         uint      `gorm:"column:user_id;primary_key;size:16;"` // yyyymmdd00000000
	UserEmail      string    `gorm:"column:user_email;size:255;unique_index;not null;"`
	UserPassword   string    `gorm:"column:user_password;size:32;not null;"`
	UserNickName   string    `gorm:"column:user_nick_name;size:64;"`
	UserImage      string    `gorm:"column:user_image;size:204800;"`
	UserStatus     uint      `gorm:"column:user_status;size:5;not null;default:1"` // 1: activated; 0: not activated
	UserCreateTime time.Time `gorm:"column:user_create_time;size:100;not null;"`
}

func (eUser User) EntityToSchemaUser() *schema.User {
	sUser := new(schema.User)
	_ = Transform(eUser, sUser)
	return sUser
}

func (eUser User) EntityToSchemaUserWithPassword() *schema.UserWithPassword {
	sUser := new(schema.UserWithPassword)
	_ = Transform(eUser, sUser)
	return sUser
}
