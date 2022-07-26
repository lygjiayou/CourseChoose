package model

// 说明：
// 1. 所提到的位数均以字节长度为准
// 2. 所有的ID均为int64(以string方式表现）

// 通用结构

type ErrNo int

const (
	OK				ErrNo = iota
	ParamInvalid	ErrNo = 1	// 参数不合法
	UserHasExisted	ErrNo = 2	// 该Username已存在
	UserHasDeleted	ErrNo = 3	// 用户已删除
	UserNotExisted	ErrNo = 4	// 用户不存在
	WrongPassword	ErrNo = 5	// 密码错误
	LoginRequired	ErrNo = 6	// 用户未登录

	UnknownError	ErrNo = 255	// 未知错误
)
// ......
// 成员管理

type UserType int

const (
	Admin 	UserType = 1 // 管理员
	Student UserType = 2 // 学生
	Teacher UserType = 3 // 教师
)

// 系统内置管理员账号
// 账号名：JudgeAdmin 密码：JudgePassword2022

// 创建成员
// 参数不合法返回 ParamInvalid

// 只有管理员才能添加
type CreateMemberRequest struct {
	Nickname string	`binding:"required"` // required(required 属性规定必需在提交表单之前填写输入字段)，不小于4位 不超过20位
	Username string	`binding:"required"` // required,必填，支持大小写，不小于8位，不超过20位（字节）
	Password string	`binding:"required"` // required,必填，同时包括大小写、数字，不少于8位不超过20位（字节）
	UserType UserType `binding:"required"` // required,必填，枚举值（1：管理员，2：学生，3：教师）
}

type CreateMemberRespone struct {
	Code ErrNo
	Data struct{
		UserID string // int64 范围
	}
}
