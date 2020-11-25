package service

import (
	"battery-analysis-platform/app/web/model"
	"battery-analysis-platform/pkg/jd"
)

type GetSysInfoService struct {
}

func (s *GetSysInfoService) Do() (*jd.Response, error) {
	data := model.NewSysInfo()
	return jd.Build(jd.SUCCESS, "", data), nil
}
