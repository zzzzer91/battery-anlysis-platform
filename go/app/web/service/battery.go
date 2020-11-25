package service

import (
	"battery-analysis-platform/app/web/constant"
	"battery-analysis-platform/app/web/dal/mongo"
	"battery-analysis-platform/pkg/jd"
	"battery-analysis-platform/pkg/jtime"
	"strings"
	"time"
)

type GetBatteryListService struct {
	DataComeFrom string `form:"dataComeFrom"`
	StartDate    string `form:"startDate"`
	NeedParams   string `form:"needParams"`
	DataLimit    int    `form:"dataLimit"`
}

func (s *GetBatteryListService) Do() (*jd.Response, error) {
	// 校验字段合法性
	table, ok := constant.BatteryNameToTable[s.DataComeFrom]
	if !ok {
		return jd.Err("参数 dataComeFrom 不合法"), nil
	}
	fields := strings.Split(s.NeedParams, ",")
	for _, field := range fields {
		_, ok = table.FieldSet[field]
		if !ok {
			return jd.Err("参数 needParams 不合法"), nil
		}
	}
	startDate, err := time.ParseInLocation(jtime.FormatLayout, s.StartDate, time.Local)
	if err != nil {
		return jd.Err("参数 startDate 不合法"), nil
	}
	if s.DataLimit > 10000 {
		return jd.Err("参数 dataLimit 不合法"), nil
	}

	data, err := mongo.GetBatteryList(table.Name, startDate, s.DataLimit, fields)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return jd.Err("未查询到相关数据"), nil
	}
	return jd.Build(jd.SUCCESS, "查询成功", data), nil
}
