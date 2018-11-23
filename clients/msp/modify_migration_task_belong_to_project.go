package msp

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 更改迁移任务所属项目
// https://cloud.tencent.com/document/api/659/18601

type ModifyMigrationTaskBelongToProjectRequest struct {
	// 项目ID
	ProjectId int64 `name:"ProjectId"`
	// 区域
	Region string `name:"Region"`
	// 任务ID
	TaskId string `name:"TaskId"`
}

func (req *ModifyMigrationTaskBelongToProjectRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyMigrationTaskBelongToProjectResponse, error) {
	resp := &ModifyMigrationTaskBelongToProjectResponse{}
	err := client.Request("msp", "ModifyMigrationTaskBelongToProject", "2018-03-19").Do(req, resp)
	return resp, err
}

type ModifyMigrationTaskBelongToProjectResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
