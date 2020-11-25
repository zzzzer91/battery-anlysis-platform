package model

import (
	"battery-analysis-platform/app/web/constant"
	"battery-analysis-platform/pkg/jtime"
	"battery-analysis-platform/pkg/security"
)

type User struct {
	Name     string `json:"userName" bson:"name"`
	Password string `json:"-" bson:"password"`
	Type     int    `json:"userType" bson:"type"`
	// *string 让其 json 时可以返回 null，否则只能返回字符串零值
	AvatarName    *string    `json:"avatarName" bson:"avatarName"`
	LastLoginTime jtime.Time `json:"lastLoginTime" bson:"lastLoginTime"`
	Comment       string     `json:"comment" bson:"comment"`
	LoginCount    int        `json:"loginCount" bson:"loginCount"`
	Status        int        `json:"userStatus" bson:"status"`
	CreateTime    jtime.Time `json:"createTime" bson:"createTime"`
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) {
	s, _ := security.GeneratePasswordHash(password)
	user.Password = s
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := security.CheckPasswordHash(user.Password, password)
	return err == nil
}

func (user *User) CheckStatusOk() bool {
	return user.Status == constant.UserStatusNormal
}

func NewUser(name, password, comment string) *User {
	now := jtime.Now()
	user := User{
		Name:          name,
		Comment:       comment,
		CreateTime:    now,
		LastLoginTime: now,
		Status:        1,
	}
	user.SetPassword(password)
	return &user
}
