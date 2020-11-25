package mongo

import (
	"battery-analysis-platform/app/web/constant"
	"battery-analysis-platform/app/web/db"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newTimeoutCtx() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), constant.MongoCtxTimeout)
	return ctx
}

// 确保创建 mongo 索引
func createMongoCollectionIdx(name string, model mongo.IndexModel) error {
	collection := db.Mongo.Collection(name)
	ctx := newTimeoutCtx()
	_, err := collection.Indexes().CreateOne(ctx, model)
	return err
}

// 在 collection 中插入一条记录
func insertMongoCollection(collectionName string, item interface{}) error {
	collection := db.Mongo.Collection(collectionName)
	ctx := newTimeoutCtx()
	_, err := collection.InsertOne(ctx, item)
	return err
}

func init() {
	// user
	indexModel := mongo.IndexModel{
		Keys: bson.M{
			"name": 1,
		},
		Options: options.Index().SetUnique(true),
	}
	if err := createMongoCollectionIdx(constant.MongoCollectionUser, indexModel); err != nil {
		panic(err)
	}
	indexModel = mongo.IndexModel{
		Keys: bson.M{
			"type": 1,
		},
		Options: options.Index().SetUnique(false),
	}
	if err := createMongoCollectionIdx(constant.MongoCollectionUser, indexModel); err != nil {
		panic(err)
	}

	// yutong_vehicle
	indexModel = mongo.IndexModel{
		Keys: bson.M{
			"时间": 1,
		},
		Options: options.Index().SetUnique(false),
	}
	if err := createMongoCollectionIdx(constant.MongoCollectionYuTongVehicle, indexModel); err != nil {
		panic(err)
	}
	indexModel = mongo.IndexModel{
		Keys: bson.M{
			"状态号": 1,
		},
		Options: options.Index().SetUnique(false),
	}
	if err := createMongoCollectionIdx(constant.MongoCollectionYuTongVehicle, indexModel); err != nil {
		panic(err)
	}

	// beiqi_vehicle
	indexModel = mongo.IndexModel{
		Keys: bson.M{
			"时间": 1,
		},
		Options: options.Index().SetUnique(false),
	}
	if err := createMongoCollectionIdx(constant.MongoCollectionBeiQiVehicle, indexModel); err != nil {
		panic(err)
	}
	indexModel = mongo.IndexModel{
		Keys: bson.M{
			"状态号": 1,
		},
		Options: options.Index().SetUnique(false),
	}
	if err := createMongoCollectionIdx(constant.MongoCollectionBeiQiVehicle, indexModel); err != nil {
		panic(err)
	}

	// task
	indexModel = mongo.IndexModel{
		Keys: bson.M{
			"taskId": 1,
		},
		Options: options.Index().SetUnique(false),
	}
	if err := createMongoCollectionIdx(constant.MongoCollectionMiningTask, indexModel); err != nil {
		panic(err)
	}
	if err := createMongoCollectionIdx(constant.MongoCollectionDlTask, indexModel); err != nil {
		panic(err)
	}
}
