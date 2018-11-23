package msp

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取迁移任务列表
// https://cloud.tencent.com/document/api/659/18602

type ListMigrationTaskRequest struct {
	// 记录条数，默认值为10
	Limit *int64 `name:"Limit,omitempty"`
	// 记录起始数，默认值为0
	Offset *int64 `name:"Offset,omitempty"`
	// 项目ID，默认值为空
	ProjectId *int64 `name:"ProjectId,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *ListMigrationTaskRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ListMigrationTaskResponse, error) {
	resp := &ListMigrationTaskResponse{}
	err := client.Request("msp", "ListMigrationTask", "2018-03-19").Do(req, resp)
	return resp, err
}

type ListMigrationTaskResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 迁移任务列表
	Tasks []*Task `json:"Tasks"`
	// 记录总条数
	TotalCount int64 `json:"TotalCount"`
}
