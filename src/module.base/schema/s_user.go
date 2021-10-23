/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/8/5 1:42
 * @version     v1.0
 * @filename    s_blog.go
 * @description
 ***************************************************************************/
package schema

import "time"

type User struct {
	UserId         uint      `json:"user_id"` // yyyymmdd00000000
	UserEmail      string    `json:"user_email"`
	UserNickName   string    `json:"user_nick_name"`
	UserImage      string    `json:"user_image"`
	UserStatus     uint      `json:"user_status"` // 1: activated; 0: not activated
	UserCreateTime time.Time `json:"user_create_time"`
}

type UserWithPassword struct {
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
}

type RegisterInfo struct {
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
}
