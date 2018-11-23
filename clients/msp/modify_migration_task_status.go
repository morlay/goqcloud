package msp

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 更新迁移任务状态
// https://cloud.tencent.com/document/api/659/18600

type ModifyMigrationTaskStatusRequest struct {
	// 区域
	Region string `name:"Region"`
	// 任务状态
	Status string `name:"Status"`
	// 任务ID
	TaskId string `name:"TaskId"`
}

func (req *ModifyMigrationTaskStatusRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyMigrationTaskStatusResponse, error) {
	resp := &ModifyMigrationTaskStatusResponse{}
	err := client.Request("msp", "ModifyMigrationTaskStatus", "2018-03-19").Do(req, resp)
	return resp, err
}

type ModifyMigrationTaskStatusResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
