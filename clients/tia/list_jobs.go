package tia

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 列举Job
// https://cloud.tencent.com/document/api/851/18305

type ListJobsRequest struct {
	// 运行任务的集群
	Cluster string `name:"Cluster"`
	// 分页参数，返回数量
	Limit *int64 `name:"Limit,omitempty"`
	// 分页参数，起始位置
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *ListJobsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ListJobsResponse, error) {
	resp := &ListJobsResponse{}
	err := client.Request("tia", "ListJobs", "2018-02-26").Do(req, resp)
	return resp, err
}

type ListJobsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 训练任务列表
	Jobs []*Job `json:"Jobs"`
}
