package batch

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查看作业列表
// https://cloud.tencent.com/document/api/599/15909
type DescribeJobsRequest struct {
	// 过滤条件
	Filters []*Filter `name:"Filters,omitempty"`
	// 作业ID
	JobIds []*string `name:"JobIds,omitempty"`
	// 返回数量
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeJobsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeJobsResponse, error) {
	resp := &DescribeJobsResponse{}
	err := client.Request("batch", "DescribeJobs", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeJobsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 作业列表
	JobSet JobView `json:"JobSet"`
	// 符合条件的作业数量
	TotalCount int64 `json:"TotalCount"`
}
