package batch

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取计算环境列表
// https://cloud.tencent.com/document/api/599/15893
type DescribeComputeEnvsRequest struct {
	// 计算环境ID
	EnvIds []*string `name:"EnvIds,omitempty"`
	// 过滤条件
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeComputeEnvsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeComputeEnvsResponse, error) {
	resp := &DescribeComputeEnvsResponse{}
	err := client.Request("batch", "DescribeComputeEnvs", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeComputeEnvsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 计算环境列表
	ComputeEnvSet []*ComputeEnvView `json:"ComputeEnvSet"`
	// 计算环境数量
	TotalCount int64 `json:"TotalCount"`
}
