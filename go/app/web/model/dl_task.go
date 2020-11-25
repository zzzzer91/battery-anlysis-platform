package model

type NnLayer struct {
	Neurons    int    `json:"neurons" bson:"neurons"`
	Activation string `json:"activation" bson:"activation"`
}

type NnHyperParameter struct {
	HiddenLayerStructure  []NnLayer `json:"hiddenLayerStructure" bson:"hiddenLayerStructure"`
	OutputLayerActivation string    `json:"outputLayerActivation" bson:"outputLayerActivation"`
	Loss                  string    `json:"loss" bson:"loss"`
	Seed                  int       `json:"seed" bson:"seed"`
	BatchSize             int       `json:"batchSize" bson:"batchSize"`
	Epochs                int       `json:"epochs" bson:"epochs"`
	LearningRate          float64   `json:"learningRate" bson:"learningRate"`
}

type NnTrainingHistory struct {
	Loss     []float64 `json:"loss" bson:"loss"`
	Accuracy []float64 `json:"accuracy" bson:"accuracy"`
}

type NnEvalResult struct {
	A1Count     int `json:"a1Count" bson:"a1Count"`
	A2Count     int `json:"a2Count" bson:"a2Count"`
	A3Count     int `json:"a3Count" bson:"a3Count"`
	A4Count     int `json:"a4Count" bson:"a4Count"`
	AOtherCount int `json:"aOtherCount" bson:"aOtherCount"`
}

type DlTask struct {
	BaseTask        `bson:",inline"`
	Dataset         string             `json:"dataset" bson:"dataset"`
	HyperParameter  *NnHyperParameter  `json:"hyperParameter" bson:"hyperParameter"`
	TrainingHistory *NnTrainingHistory `json:"-" bson:"trainingHistory"`
	EvalResult      *NnEvalResult      `json:"-" bson:"evalResult"`
}

func NewDlTask(id, dataset string, hyperParameter *NnHyperParameter) *DlTask {
	return &DlTask{
		BaseTask:       NewBaseTask(id),
		Dataset:        dataset,
		HyperParameter: hyperParameter,
	}
}
