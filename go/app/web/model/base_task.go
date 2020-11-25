package model

import (
	"battery-analysis-platform/app/web/constant"
	"battery-analysis-platform/pkg/jtime"
)

// 注意：
// 结构体嵌入其他结构体时，要指定 `bson:",inline"`（注意有个逗号）标签才能正确读入 bson，
// 同时该结构体要public；
// 而用 json 序列化时，直接嵌入结构体即可，不需要做其他事。
// 嵌入时，不能用指针
type BaseTask struct {
	TaskId     string     `json:"taskId" bson:"taskId"`
	CreateTime jtime.Time `json:"createTime" bson:"createTime"`
	TaskStatus int        `json:"taskStatus" bson:"taskStatus"`
	Comment    string     `json:"comment" bson:"comment"`
}

func NewBaseTask(id string) BaseTask {
	return BaseTask{
		TaskId:     id,
		CreateTime: jtime.Now(),
		TaskStatus: constant.TaskStatusPreparing,
	}
}
