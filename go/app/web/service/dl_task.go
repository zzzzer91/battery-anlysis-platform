package service

import (
	"battery-analysis-platform/app/web/conf"
	"battery-analysis-platform/app/web/constant"
	"battery-analysis-platform/app/web/dal/mongo"
	"battery-analysis-platform/app/web/db"
	"battery-analysis-platform/app/web/model"
	"battery-analysis-platform/app/web/producer"
	"battery-analysis-platform/pkg/jd"
	"fmt"
)

type CreateDlTaskService struct {
	Dataset        string                  `json:"dataset"`
	HyperParameter *model.NnHyperParameter `json:"hyperParameter"`
}

func (s *CreateDlTaskService) Do() (*jd.Response, error) {
	// TODO 检查输入参数

	// 检查是否达到创建任务上限
	if !producer.CheckTaskLimit(constant.RedisKeyDlTaskWorkingIdSet, 1) {
		return jd.Err("允许同时执行任务数已达上限"), nil
	}

	asyncResult, err := producer.Celery.Delay(
		constant.CeleryTaskDeeplearningTrain, s.Dataset, s.HyperParameter)
	if err != nil {
		return nil, err
	}
	// 添加正在工作的任务的 id 到集合中
	err = producer.AddWorkingTaskIdToSet(constant.RedisKeyDlTaskWorkingIdSet, asyncResult.TaskID)
	if err != nil {
		return nil, err
	}

	data, err := mongo.CreateDlTask(asyncResult.TaskID, s.Dataset, s.HyperParameter)
	if err != nil {
		return nil, err
	}

	return jd.Build(jd.SUCCESS, "创建成功", data), nil
}

type GetDlTaskListService struct {
}

func (s *GetDlTaskListService) Do() (*jd.Response, error) {
	data, err := mongo.GetDlTaskList()
	if err != nil {
		return nil, err
	}
	return jd.Build(jd.SUCCESS, "", data), nil
}

type GetDlTaskTraningHistoryService struct {
	Id            string
	ReadFromRedis bool
}

func (s *GetDlTaskTraningHistoryService) Do() (*jd.Response, error) {
	data, err := mongo.GetDlTaskTrainingHistory(s.Id, s.ReadFromRedis)
	if err != nil {
		return nil, err
	}
	return jd.Build(jd.SUCCESS, "", data), nil
}

type GetDlTaskEvalResultService struct {
	Id string
}

func (s *GetDlTaskEvalResultService) Do() (*jd.Response, error) {
	data, err := mongo.GetDlTaskEvalResult(s.Id)
	if err != nil {
		return nil, err
	}
	return jd.Build(jd.SUCCESS, "", data), nil
}

type DownloadDlModelService struct {
	Id string
}

func (s *DownloadDlModelService) Do() (string, error) {
	return conf.App.Gin.ResourcePath + constant.FileDlModelPath + fmt.Sprintf("/%s.pt", s.Id), nil
}

type DeleteDlTaskService struct {
	Id string
}

func (s *DeleteDlTaskService) Do() (*jd.Response, error) {
	// 因为 gocelery 未提供终止任务的 api，这里把终止行为封装成任务，然后调用它
	_, err := producer.Celery.Delay(constant.CeleryTaskDeeplearningStopTrain, s.Id)
	if err != nil {
		return nil, err
	}

	err = producer.DelWorkingTaskIdFromSet(constant.RedisKeyDlTaskWorkingIdSet, s.Id)
	if err != nil {
		return nil, err
	}

	// 删除暂存在 redis 中的数据
	prefixStr := constant.RedisPrefixDlTaskTrainingHistory + s.Id + ":"
	db.Redis.Del(prefixStr+constant.RedisCommonKeySigList, prefixStr+"loss", prefixStr+"accuracy")

	err = mongo.DeleteDlTask(s.Id)
	if err != nil {
		return nil, err
	}

	return jd.Build(jd.SUCCESS, "删除成功", nil), nil
}
