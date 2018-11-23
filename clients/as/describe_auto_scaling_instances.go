package as

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询实例
// https://cloud.tencent.com/document/api/377/20437

type DescribeAutoScalingInstancesRequest struct {
	// 过滤条件。 instance-id - String - 是否必填：否 -（过滤条件）按照实例ID过滤。 auto-scaling-group-id - String - 是否必填：否 -（过滤条件）按照伸缩组ID过滤。每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定InstanceIds和Filters。
	Filters []*Filter `name:"Filters,omitempty"`
	// 待查询的云主机（CVM）实例ID。参数不支持同时指定InstanceIds和Filters。
	InstanceIds []*string `name:"InstanceIds,omitempty"`
	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeAutoScalingInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeAutoScalingInstancesResponse, error) {
	resp := &DescribeAutoScalingInstancesResponse{}
	err := client.Request("as", "DescribeAutoScalingInstances", "2018-04-19").Do(req, resp)
	return resp, err
}

type DescribeAutoScalingInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 实例详细信息列表。
	AutoScalingInstanceSet []*Instance `json:"AutoScalingInstanceSet"`
	// 符合条件的实例数量。
	TotalCount int64 `json:"TotalCount"`
}
