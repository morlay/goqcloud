package msp

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 取消注册迁移任务
// https://cloud.tencent.com/document/api/659/18605

type DeregisterMigrationTaskRequest struct {
	// 区域
	Region string `name:"Region"`
	// 任务ID
	TaskId string `name:"TaskId"`
}

func (req *DeregisterMigrationTaskRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeregisterMigrationTaskResponse, error) {
	resp := &DeregisterMigrationTaskResponse{}
	err := client.Request("msp", "DeregisterMigrationTask", "2018-03-19").Do(req, resp)
	return resp, err
}

type DeregisterMigrationTaskResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
