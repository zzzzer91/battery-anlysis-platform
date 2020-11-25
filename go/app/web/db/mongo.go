package db

import (
	"battery-analysis-platform/app/web/conf"
	"battery-analysis-platform/pkg/db"
	"go.mongodb.org/mongo-driver/mongo"
)

var Mongo *mongo.Database

func init() {
	d, err := db.InitMongo(&conf.App.Mongo)
	if err != nil {
		panic(err)
	}
	Mongo = d
}
