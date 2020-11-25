package mongo

import (
	"battery-analysis-platform/app/web/constant"
	"battery-analysis-platform/app/web/db"
	"battery-analysis-platform/app/web/model"
	"battery-analysis-platform/pkg/conv"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateDlTask(id, dataset string, hyperParameter *model.NnHyperParameter) (*model.DlTask, error) {
	task := model.NewDlTask(id, dataset, hyperParameter)
	err := creatTask(constant.MongoCollectionDlTask, task)

	return task, err
}

func GetDlTaskList() ([]model.DlTask, error) {
	collection := db.Mongo.Collection(constant.MongoCollectionDlTask)
	filter := bson.D{}
	projection := bson.D{
		{"_id", false},
		{"trainingHistory", false},
		{"evalResult", false},
	}
	sort := bson.D{{"createTime", -1}}
	// 注意 ctx 不能几个连接复用
	ctx, _ := context.WithTimeout(context.Background(), constant.MongoCtxTimeout)
	cur, err := collection.Find(ctx, filter, options.Find().SetProjection(projection).SetSort(sort))
	if err != nil {
		return nil, err
	}
	// 为了使其找不到时返回空列表，而不是 nil
	records := make([]model.DlTask, 0)
	ctx, _ = context.WithTimeout(context.Background(), constant.MongoCtxTimeout)
	for cur.Next(ctx) {
		result := model.DlTask{}
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		records = append(records, result)
	}
	_ = cur.Close(ctx)
	return records, nil
}

func GetDlTaskTrainingHistory(id string, readFromRedis bool) (*model.NnTrainingHistory, error) {
	if readFromRedis {
		prefixStr := constant.RedisPrefixDlTaskTrainingHistory + id + ":"

		lossStrList, err := db.Redis.LRange(
			prefixStr+"loss", 0, -1).Result()
		if err != nil {
			return nil, err
		}
		// 转换为 float
		lossList, err := conv.StringSlice2FloatSlice(lossStrList)
		if err != nil {
			return nil, err
		}

		accuracyStrList, err := db.Redis.LRange(
			prefixStr+"accuracy", 0, -1).Result()
		if err != nil {
			return nil, err
		}
		accuracyList, err := conv.StringSlice2FloatSlice(accuracyStrList)
		if err != nil {
			return nil, err
		}

		return &model.NnTrainingHistory{
			Loss:     lossList,
			Accuracy: accuracyList,
		}, nil
	} else {
		collection := db.Mongo.Collection(constant.MongoCollectionDlTask)
		filter := bson.D{{"taskId", id}}
		projection := bson.D{{"_id", false}, {"trainingHistory", true}}
		var result model.DlTask
		ctx := newTimeoutCtx()
		err := collection.FindOne(ctx, filter,
			options.FindOne().SetProjection(projection)).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result.TrainingHistory, nil
	}
}

func GetDlTaskEvalResult(id string) (*model.NnEvalResult, error) {
	collection := db.Mongo.Collection(constant.MongoCollectionDlTask)
	filter := bson.D{{"taskId", id}}
	projection := bson.D{{"_id", false}, {"evalResult", true}}
	var result model.DlTask
	ctx := newTimeoutCtx()
	err := collection.FindOne(ctx, filter,
		options.FindOne().SetProjection(projection)).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result.EvalResult, nil
}

func DeleteDlTask(id string) error {
	return deleteTask(constant.MongoCollectionDlTask, id)
}
