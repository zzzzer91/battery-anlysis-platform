package service

import (
	"battery-analysis-platform/app/web/constant"
	"battery-analysis-platform/app/web/dal/mongo"
	"battery-analysis-platform/pkg/checker"
	"battery-analysis-platform/pkg/jd"
)

type CreateUserService struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Comment  string `json:"comment"`
}

func (s *CreateUserService) Do() (*jd.Response, error) {
	if !checker.ReUserNameOrPassword.MatchString(s.UserName) {
		return jd.Err("用户名不符合要求"), nil
	}
	if !checker.ReUserNameOrPassword.MatchString(s.Password) {
		return jd.Err("密码不符合要求"), nil
	}
	if len(s.Comment) > 64 {
		return jd.Err("参数 comment 不合法"), nil
	}

	data, err := mongo.CreateUser(s.UserName, s.Password, s.Comment)
	if err != nil {
		return jd.Err("用户 " + s.UserName + " 已存在"), nil
	}
	return jd.Build(jd.SUCCESS, "创建用户 "+s.UserName+" 成功", data), nil
}

type GetCommonUserListService struct {
}

func (s *GetCommonUserListService) Do() (*jd.Response, error) {
	userList, err := mongo.GetCommonUserList()
	if err != nil {
		return nil, err
	}
	return jd.Build(jd.SUCCESS, "", userList), nil
}

type UpdateUserInfoService struct {
	UserName string
	Comment  string `json:"comment"`
	Status   int    `json:"userStatus"`
}

func (s *UpdateUserInfoService) Do() (*jd.Response, error) {
	if s.Status != constant.UserStatusNormal && s.Status != constant.UserStatusForbiddenLogin {
		return jd.Err("参数 status 不合法"), nil
	}
	if len(s.Comment) > 64 {
		return jd.Err("参数 comment 不合法"), nil
	}

	user, err := mongo.GetUser(s.UserName)
	if err != nil {
		return nil, err
	}
	user.Comment = s.Comment
	user.Status = s.Status
	err = mongo.UpdateUserInfo(user)
	if err != nil {
		return nil, err
	}
	// 用户数据发生修改，更新将其缓存
	err = mongo.AddUserToCache(user)
	if err != nil {
		return nil, err
	}
	return jd.Build(jd.SUCCESS, "修改用户 "+s.UserName+" 成功", user), nil
}

type UpdateUserPasswordService struct {
	UserName string
	Password string `json:"password"`
}

func (s *UpdateUserPasswordService) Do() (*jd.Response, error) {
	err := mongo.UpdateUserPassword(s.UserName, s.Password)
	if err != nil {
		return nil, err
	}
	return jd.Build(jd.SUCCESS, "修改密码成功", nil), nil
}
