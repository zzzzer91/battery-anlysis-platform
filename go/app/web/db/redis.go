package db

import (
	"battery-analysis-platform/app/web/conf"
	"battery-analysis-platform/pkg/db"
	"github.com/go-redis/redis/v7"
)

var Redis *redis.Client

func init() {
	d, err := db.InitRedis(&conf.App.Redis)
	if err != nil {
		panic(err)
	}
	Redis = d
}
