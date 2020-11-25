package constant

// 用户类型 Type
const (
	// 超级用户
	UserTypeSuperUser = 64
	// 普通用户
	UserTypeCommonUser = 0
)

// 用户状态 Status
const (
	UserStatusForbiddenLogin = 0
	UserStatusNormal         = 1
)

// 存在 cookie 中的 key 名
const (
	CookieKey = "userName"
)
