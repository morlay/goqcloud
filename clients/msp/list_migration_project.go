package msp

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取迁移项目名称列表
// https://cloud.tencent.com/document/api/659/18603

type ListMigrationProjectRequest struct {
	// 返回条数，默认值为500
	Limit *int64 `name:"Limit,omitempty"`
	// 记录起始数，默认值为0
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *ListMigrationProjectRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ListMigrationProjectResponse, error) {
	resp := &ListMigrationProjectResponse{}
	err := client.Request("msp", "ListMigrationProject", "2018-03-19").Do(req, resp)
	return resp, err
}

type ListMigrationProjectResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 项目列表
	Projects []*Project `json:"Projects"`
	// 项目总数
	TotalCount int64 `json:"TotalCount"`
}
