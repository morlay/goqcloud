package iot

import (
	time "time"
)

// 用户
type User struct {
	// app_id
	AppId int64 `json:"AppId"`
	// 用户类型（1：国内，2：国际）
	Area int64 `json:"Area"`
	// 计费类型（日结、月结）
	BillingType string `json:"BillingType"`
	// 创建时间
	CreateTime time.Time `json:"CreateTime"`
	// 备注信息
	Message string `json:"Message"`
	// 用户状态（0：正常，1：欠费，2：恶意）
	Status int64 `json:"Status"`
	// 修改时间
	UpdateTime time.Time `json:"UpdateTime"`
}

// 设备
type Device struct {
	// 创建时间
	CreateTime time.Time `json:"CreateTime"`
	// 设备信息
	DeviceInfo Object `json:"DeviceInfo"`
	// 设备名称
	DeviceName string `json:"DeviceName"`
	// 设备密钥
	DeviceSecret string `json:"DeviceSecret"`
	// 产品Id
	ProductId string `json:"ProductId"`
	// 更新时间
	UpdateTime time.Time `json:"UpdateTime"`
}

// 对象
type Object *Object // 规则
type Rule struct {
	// 转发
	Actions []*Object `json:"Actions"`
	// 已启动
	Active int64 `json:"Active"`
	// AppId
	AppId int64 `json:"AppId"`
	// 创建时间
	CreateTime time.Time `json:"CreateTime"`
	// 已删除
	Deleted int64 `json:"Deleted"`
	// 描述
	Description string `json:"Description"`
	// 名称
	Name string `json:"Name"`
	// 查询
	Query RuleQuery `json:"Query"`
	// 规则Id
	RuleId string `json:"RuleId"`
	// 更新时间
	UpdateTime time.Time `json:"UpdateTime"`
}

// 查询
type RuleQuery struct {
	// 过滤规则
	Condition string `json:"Condition"`
	// 字段
	Field string `json:"Field"`
	// Topic
	Topic string `json:"Topic"`
}

// Topic
type Topic struct {
	// 创建时间
	CreateTime time.Time `json:"CreateTime"`
	// 已删除
	Deleted int64 `json:"Deleted"`
	// 消息最大数量
	MsgCount int64 `json:"MsgCount"`
	// 消息最大生命周期
	MsgLife int64 `json:"MsgLife"`
	// 消息最大大小
	MsgSize int64 `json:"MsgSize"`
	// Topic完整路径
	Path string `json:"Path"`
	// 产品Id
	ProductId string `json:"ProductId"`
	// TopicId
	TopicId string `json:"TopicId"`
	// Topic名称
	TopicName string `json:"TopicName"`
	// 更新时间
	UpdateTime `json:"UpdateTime"`
}

// 应用用户
type AppUser struct {
	// 应用Id
	ApplicationId string `json:"ApplicationId"`
	// 创建时间
	CreateTime `json:"CreateTime"`
	// 绑定设备列表
	Devices []*Object `json:"Devices"`
	// 昵称
	NickName string `json:"NickName"`
	// 修改时间
	UpdateTime `json:"UpdateTime"`
	// 用户名
	UserName string `json:"UserName"`
}

// 设备状态
type DeviceStatus struct {
	// 设备名称
	DeviceName string `json:"DeviceName"`
	// 设备状态（inactive, online, offline）
	Status string `json:"Status"`
}

// 产品
type Product struct {
	// AppId
	AppId int64 `json:"AppId"`
	// 鉴权类型（0：直连，1：Token）
	AuthType int64 `json:"AuthType"`
	// 创建时间
	CreateTime time.Time `json:"CreateTime"`
	// 数据模版
	DataTemplate Object `json:"DataTemplate"`
	// 删除（0未删除）
	Deleted int64 `json:"Deleted"`
	// 产品描述
	Description string `json:"Description"`
	// 连接域名
	Domain string `json:"Domain"`
	// 备注
	Message string `json:"Message"`
	// 产品名称
	Name string `json:"Name"`
	// 产品Id
	ProductId string `json:"ProductId"`
	// 产品Key
	ProductKey string `json:"ProductKey"`
	// 产品直连密钥
	ProductSecret string `json:"ProductSecret"`
	// 产品规格
	Standard int64 `json:"Standard"`
	// 更新时间
	UpdateTime time.Time `json:"UpdateTime"`
}
