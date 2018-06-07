package batch

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查看计算环境创建信息列表
// https://cloud.tencent.com/document/api/599/15894

type DescribeComputeEnvCreateInfosRequest struct {
	// 计算环境ID
	EnvIds []*string `name:"EnvIds,omitempty"`
	// 过滤条件 zone - String - 是否必填：否 -（过滤条件）按照可用区过滤。 env-id - String - 是否必填：否 -（过滤条件）按照计算环境ID过滤。 env-name - String - 是否必填：否 -（过滤条件）按照计算环境名称过滤。
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeComputeEnvCreateInfosRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeComputeEnvCreateInfosResponse, error) {
	resp := &DescribeComputeEnvCreateInfosResponse{}
	err := client.Request("batch", "DescribeComputeEnvCreateInfos", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeComputeEnvCreateInfosResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 计算环境创建信息列表
	ComputeEnvCreateInfoSet []*ComputeEnvCreateInfo `json:"ComputeEnvCreateInfoSet"`
	// 计算环境数量
	TotalCount int64 `json:"TotalCount"`
}
