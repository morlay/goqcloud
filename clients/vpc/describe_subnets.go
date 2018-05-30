package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询子网列表
// https://cloud.tencent.com/document/api/215/15784
type DescribeSubnetsRequest struct {
	// 过滤条件，参数不支持同时指定SubnetIds和Filters。subnet-id - String - （过滤条件）Subnet实例名称。vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。cidr-block - String - （过滤条件）vpc的cidr。is-default - Boolean - （过滤条件）是否是默认子网。subnet-name - String - （过滤条件）子网名称。zone - String - （过滤条件）可用区。
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量
	Limit *string `name:"Limit,omitempty"`
	// 偏移量
	Offset *string `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 子网实例ID查询。形如：subnet-pxir56ns。每次请求的实例的上限为100。参数不支持同时指定SubnetIds和Filters。
	SubnetIds []*string `name:"SubnetIds,omitempty"`
}

func (req *DescribeSubnetsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeSubnetsResponse, error) {
	resp := &DescribeSubnetsResponse{}
	err := client.Request("vpc", "DescribeSubnets", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeSubnetsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 子网对象。
	SubnetSet []*Subnet `json:"SubnetSet"`
	// 符合条件的实例数量。
	TotalCount int64 `json:"TotalCount"`
}
