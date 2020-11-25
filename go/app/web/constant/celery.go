package constant

// 可以调用的 celery 任务名
const (
	CeleryTaskMiningComputeModel     = "task.mining.compute_model"
	CeleryTaskMiningStopComputeModel = "task.mining.stop_compute_model"

	CeleryTaskDeeplearningTrain     = "task.deeplearning.train"
	CeleryTaskDeeplearningStopTrain = "task.deeplearning.stop_train"
)
