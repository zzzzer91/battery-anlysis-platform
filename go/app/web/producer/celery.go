package producer

import (
	"battery-analysis-platform/app/web/conf"
	"battery-analysis-platform/app/web/db"
	"battery-analysis-platform/pkg/producer"
	"github.com/gocelery/gocelery"
)

var Celery *gocelery.CeleryClient

func init() {
	cli, err := producer.InitCelery(&conf.App.Celery)
	if err != nil {
		panic(err)
	}
	Celery = cli
}

func CheckTaskLimit(key string, limit int) bool {
	return producer.CheckTaskLimit(db.Redis, key, limit)
}

func AddWorkingTaskIdToSet(key string, id string) error {
	return producer.AddWorkingTaskIdToSet(db.Redis, key, id)
}

func DelWorkingTaskIdFromSet(key string, id string) error {
	return producer.DelWorkingTaskIdFromSet(db.Redis, key, id)
}
