package model

// 不包含任务数据，这个 model 用来做任务列表的元素，所以不需要
type MiningTask struct {
	BaseTask     `bson:",inline"`
	TaskName     string `json:"taskName" bson:"taskName"`
	DataComeFrom string `json:"dataComeFrom" bson:"dataComeFrom"`
	DateRange    string `json:"dateRange" bson:"dateRange"`
}

func NewMiningTask(id, name, dataComeFrom, dateRange string) *MiningTask {
	return &MiningTask{
		BaseTask:     NewBaseTask(id),
		TaskName:     name,
		DataComeFrom: dataComeFrom,
		DateRange:    dateRange,
	}
}
