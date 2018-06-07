package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询基础网络互通列表
// https://cloud.tencent.com/document/api/215/15776

type DescribeClassicLinkInstancesRequest struct {
	// 过滤条件。vpc-id - String - （过滤条件）VPC实例ID。vm-ip - String - （过滤条件）基础网络云主机IP。
	Filters []*FilterObject `name:"Filters,omitempty"`
	// 返回数量
	Limit *string `name:"Limit,omitempty"`
	// 偏移量
	Offset *string `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeClassicLinkInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeClassicLinkInstancesResponse, error) {
	resp := &DescribeClassicLinkInstancesResponse{}
	err := client.Request("vpc", "DescribeClassicLinkInstances", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeClassicLinkInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 私有网络和基础网络互通设备。
	ClassicLinkInstanceSet []*ClassicLinkInstance `json:"ClassicLinkInstanceSet"`
	// 符合条件的实例数量。
	TotalCount int64 `json:"TotalCount"`
}
