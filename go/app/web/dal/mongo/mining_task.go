package mongo

import (
	"battery-analysis-platform/app/web/constant"
	"battery-analysis-platform/app/web/db"
	"battery-analysis-platform/app/web/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateMiningTask(id, name, dataComeFrom, dateRange string) (*model.MiningTask, error) {
	task := model.NewMiningTask(id, name, dataComeFrom, dateRange)
	err := creatTask(constant.MongoCollectionMiningTask, task)
	return task, err
}

func GetMiningTaskList() ([]model.MiningTask, error) {
	collection := db.Mongo.Collection(constant.MongoCollectionMiningTask)
	filter := bson.D{}
	projection := bson.D{{"_id", false}, {"data", false}}
	sort := bson.D{{"createTime", -1}}
	// 注意 ctx 不能几个连接复用，原因见 `context.WithTimeout` 源码
	ctx := newTimeoutCtx()
	cur, err := collection.Find(ctx, filter,
		options.Find().SetProjection(projection).SetSort(sort))
	if err != nil {
		return nil, err
	}
	// 为了使其找不到时返回空列表，而不是 nil
	records := make([]model.MiningTask, 0)
	for cur.Next(ctx) {
		result := model.MiningTask{}
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		records = append(records, result)
	}
	_ = cur.Close(ctx)
	return records, nil
}

func GetMiningTaskData(id string) (bson.A, error) {
	collection := db.Mongo.Collection(constant.MongoCollectionMiningTask)
	filter := bson.D{{"taskId", id}}
	projection := bson.D{{"_id", false}, {"data", true}}
	ctx := newTimeoutCtx()
	var result bson.M
	err := collection.FindOne(ctx, filter,
		options.FindOne().SetProjection(projection)).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result["data"].(bson.A), nil
}

func DeleteMiningTask(id string) error {
	return deleteTask(constant.MongoCollectionMiningTask, id)
}
