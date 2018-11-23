package msp

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 注册迁移任务
// https://cloud.tencent.com/document/api/659/18599

type RegisterMigrationTaskRequest struct {
	// 迁移任务创建时间
	CreateTime time.Time `name:"CreateTime"`
	// 目标实例接入类型
	DstAccessType *string `name:"DstAccessType,omitempty"`
	// 目标实例数据库类型
	DstDatabaseType *string `name:"DstDatabaseType,omitempty"`
	// 迁移任务目的信息
	DstInfo DstInfo `name:"DstInfo"`
	// 迁移类别，如数据库迁移中mysql:mysql代表从mysql迁移到mysql，文件迁移中oss:cos代表从阿里云oss迁移到腾讯云cos
	MigrateClass string `name:"MigrateClass"`
	// 区域
	Region string `name:"Region"`
	// 服务提供商名称
	ServiceSupplier string `name:"ServiceSupplier"`
	// 源实例接入类型
	SrcAccessType *string `name:"SrcAccessType,omitempty"`
	// 源实例数据库类型
	SrcDatabaseType *string `name:"SrcDatabaseType,omitempty"`
	// 迁移任务源信息
	SrcInfo SrcInfo `name:"SrcInfo"`
	// 任务名称
	TaskName string `name:"TaskName"`
	// 任务类型，取值database（数据库迁移）、file（文件迁移）、host（主机迁移）
	TaskType string `name:"TaskType"`
	// 迁移任务更新时间
	UpdateTime time.Time `name:"UpdateTime"`
}

func (req *RegisterMigrationTaskRequest) Invoke(client github_com_morlay_goqcloud.Client) (*RegisterMigrationTaskResponse, error) {
	resp := &RegisterMigrationTaskResponse{}
	err := client.Request("msp", "RegisterMigrationTask", "2018-03-19").Do(req, resp)
	return resp, err
}

type RegisterMigrationTaskResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 任务ID
	TaskId string `json:"TaskId"`
}
