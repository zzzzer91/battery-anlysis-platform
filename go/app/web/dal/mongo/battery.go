package mongo

import (
	"battery-analysis-platform/app/web/db"
	"battery-analysis-platform/pkg/jtime"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func GetBatteryList(tableName string, startDate time.Time, limit int, fields []string) ([]bson.M, error) {
	collection := db.Mongo.Collection(tableName)
	filter := bson.D{{"时间", bson.D{{"$gte", startDate}}}}
	// 注意 _id 默认会返回，需要手动过滤
	projection := bson.D{
		{"_id", false},
		{"时间", true},
		{"状态号", true},
	}
	for _, field := range fields {
		projection = append(projection, bson.E{Key: field, Value: true})
	}

	ctx := newTimeoutCtx()

	cur, err := collection.Find(ctx, filter,
		options.Find().SetProjection(projection).SetLimit(int64(limit)))
	if err != nil {
		return nil, err
	}

	// 为了使其找不到时返回空列表，而不是 nil
	records := make([]bson.M, 0)
	for cur.Next(ctx) {
		result := make(bson.M)
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		// 为了使 json 序列化时得到想要格式
		temp := result["时间"].(primitive.DateTime)
		result["时间"] = jtime.Wrap(temp.Time())
		records = append(records, result)
	}
	_ = cur.Close(ctx)
	return records, nil
}
