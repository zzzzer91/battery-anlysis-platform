package model

import (
	"battery-analysis-platform/app/web/constant"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser(t *testing.T) {
	ast := assert.New(t)
	user := &User{
		Name:   "xiaoming",
		Type:   constant.UserTypeSuperUser,
		Status: 0,
	}
	user.SetPassword("123456")

	ast.Equal(true, user.CheckPassword("123456"))
	ast.Equal(false, user.CheckStatusOk())

	expected := `{"userName":"xiaoming","userType":64,"avatarName":null,"lastLoginTime":null,"comment":"","loginCount":0,"userStatus":0,"createTime":null}`
	ret, err := json.Marshal(user)
	ast.Nil(err)
	ast.Equal(expected, string(ret))
}
