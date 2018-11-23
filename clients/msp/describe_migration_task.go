package msp

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取指定迁移任务详情
// https://cloud.tencent.com/document/api/659/18604

type DescribeMigrationTaskRequest struct {
	// 区域
	Region string `name:"Region"`
	// 任务ID
	TaskId string `name:"TaskId"`
}

func (req *DescribeMigrationTaskRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeMigrationTaskResponse, error) {
	resp := &DescribeMigrationTaskResponse{}
	err := client.Request("msp", "DescribeMigrationTask", "2018-03-19").Do(req, resp)
	return resp, err
}

type DescribeMigrationTaskResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 迁移详情列表
	TaskStatus []*TaskStatus `json:"TaskStatus"`
}
