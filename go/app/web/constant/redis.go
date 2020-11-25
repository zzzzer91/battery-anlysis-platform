package constant

import "time"

// redis 一些公共 key 名
const (
	RedisCommonKeyWorkingIdSet = "workingIdSet"
	RedisCommonKeySigList      = "sigList"
)

// redis key 前缀
const (
	RedisPrefixRoot                  = "ubattery:"
	RedisPrefixUser                  = RedisPrefixRoot + "users:"
	RedisPrefixMiningTask            = RedisPrefixRoot + "miningTask:"
	RedisPrefixDlTask                = RedisPrefixRoot + "deeplearningTask:"
	RedisPrefixDlTaskTrainingHistory = RedisPrefixDlTask + "trainingHistory:"
)

// redis key
const (
	RedisKeyMiningTaskWorkingIdSet = RedisPrefixMiningTask + RedisCommonKeyWorkingIdSet
	RedisKeyDlTaskWorkingIdSet     = RedisPrefixDlTask + RedisCommonKeyWorkingIdSet
	RedisKeyMiningTaskSigList      = RedisPrefixMiningTask + RedisCommonKeySigList
	RedisKeyDlTaskSigList          = RedisPrefixDlTask + RedisCommonKeySigList
)

const (
	// redis 中用户登录状态过期时间
	RedisExpirationUserLogin = time.Hour * 3
)
