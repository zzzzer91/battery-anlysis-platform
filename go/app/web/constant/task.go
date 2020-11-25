package constant

// 任务状态
const (
	TaskStatusPreparing  = 0
	TaskStatusProcessing = 1
	TaskStatusSuccess    = 6
	TaskStatusFailure    = 7
)

// 过滤不合法任务
var MiningSupportTaskSet = map[string]struct{}{
	"充电过程":        {},
	"工况":          {},
	"电池统计":        {},
	"pearson相关系数": {},
}
