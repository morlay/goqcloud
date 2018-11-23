package cis

import (
	time "time"
)

// 容器状态

type ContainerState struct {
	ExitCode   int64     `json:"ExitCode"`
	FinishTime time.Time `json:"FinishTime"`
	Reason     string    `json:"Reason"`
	StartTime  time.Time `json:"StartTime"`
	State      string    `json:"State"`
}

// 容器环境变量

type EnvironmentVar struct {
	// 环境变量名
	Name string `json:"Name"`
	// 环境变量值
	Value string `json:"Value"`
}

// 容器实例事件

type Event struct {
	Count     string `json:"Count"`
	FirstSeen string `json:"FirstSeen"`
	LastSeen  string `json:"LastSeen"`
	Level     string `json:"Level"`
	Message   string `json:"Message"`
	Reason    string `json:"Reason"`
}

// 过滤条件

type Filter struct {
	// 过滤字段，可选值 - Zone，VpcId，InstanceName
	Name string `json:"Name"`
	// 过滤值列表
	ValueList []*string `json:"ValueList"`
}

// 价格

type Price struct {
	DiscountPrice float64 `json:"DiscountPrice"`
	OriginalPrice float64 `json:"OriginalPrice"`
}

// 容器实例中容器结构体

type Container struct {
	// 容器启动参数
	Args []*string `json:"Args,omitempty"`
	// 容器启动命令
	Command *string `json:"Command,omitempty"`
	// 容器ID
	ContainerId *string `json:"ContainerId,omitempty"`
	// CPU，单位：核
	Cpu float64 `json:"Cpu"`
	// 当前状态
	CurrentState *ContainerState `json:"CurrentState,omitempty"`
	// 容器环境变量
	EnvironmentVars []*EnvironmentVar `json:"EnvironmentVars,omitempty"`
	// 镜像
	Image string `json:"Image"`
	// 内存，单位：Gi
	Memory float64 `json:"Memory"`
	// 容器名，由小写字母、数字和 - 组成，由小写字母开头，小写字母或数字结尾，且长度不超过 63个字符
	Name string `json:"Name"`
	// 上一次状态
	PreviousState *ContainerState `json:"PreviousState,omitempty"`
	// 重启次数
	RestartCount *int64 `json:"RestartCount,omitempty"`
	// 容器工作目录
	WorkingDir *string `json:"WorkingDir,omitempty"`
}

// 容器实例的具体信息

type ContainerInstance struct {
	// 容器列表
	Containers []*Container `json:"Containers"`
	// 创建时间
	CreateTime *time.Time `json:"CreateTime,omitempty"`
	// 容器实例ID
	InstanceId *string `json:"InstanceId,omitempty"`
	// 容器实例名称
	InstanceName string `json:"InstanceName"`
	// 内网IP
	LanIp *string `json:"LanIp,omitempty"`
	// 重启策略
	RestartPolicy string `json:"RestartPolicy"`
	// 启动时间
	StartTime *time.Time `json:"StartTime,omitempty"`
	// 容器实例状态
	State *string `json:"State,omitempty"`
	// 子网Cidr
	SubnetCidr *string `json:"SubnetCidr,omitempty"`
	// 容器实例所属SubnetId
	SubnetId string `json:"SubnetId"`
	// SubnetName
	SubnetName *string `json:"SubnetName,omitempty"`
	// VpcCidr
	VpcCidr *string `json:"VpcCidr,omitempty"`
	// 容器实例所属VpcId
	VpcId string `json:"VpcId"`
	// Vpc名称
	VpcName *string `json:"VpcName,omitempty"`
	// 可用区
	Zone string `json:"Zone"`
}

// 容器日志

type ContainerLog struct {
	Log  string `json:"Log"`
	Name string `json:"Name"`
	Time string `json:"Time"`
}
