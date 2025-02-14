package handler

// 人机校验
const Answer = "Ice"

// 性别
var Gender = map[int8]struct{}{
	1: {}, // 男
	2: {}, // 女
	3: {}, // 未知
}

// 申请状态
const (
	Pending  int8 = iota + 1 // 待处理
	Accepted                 // 已同意
	Rejected                 // 已拒绝
	Ignored                  // 已忽略
)

// 角色
const (
	Owner  int8 = iota + 1 // 创建者
	Admin                  // 管理员
	Member                 // 普通成员
)
