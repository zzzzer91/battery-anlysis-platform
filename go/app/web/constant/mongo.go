package constant

import "time"

const (
	MongoCtxTimeout = time.Second * 10
)

const (
	MongoCollectionUser          = "user"
	MongoCollectionYuTongVehicle = "yutong_vehicle"
	MongoCollectionBeiQiVehicle  = "beiqi_vehicle"
	MongoCollectionMiningTask    = "mining_task"
	MongoCollectionDlTask        = "deeplearning_task"
)
