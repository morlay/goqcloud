package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询VPC列表
// https://cloud.tencent.com/document/api/215/15778
type DescribeVpcsRequest struct {
	// 过滤条件，参数不支持同时指定VpcIds和Filters。vpc-name - String - （过滤条件）VPC实例名称。is-default - Boolean - （过滤条件）是否默认VPC。vpc-id - String - （过滤条件）VPC实例ID形如：vpc-f49l6u0z。cidr-block - String - （过滤条件）vpc的cidr。
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量
	Limit *string `name:"Limit,omitempty"`
	// 偏移量
	Offset *string `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// VPC实例ID。形如：vpc-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpcIds和Filters。
	VpcIds []*string `name:"VpcIds,omitempty"`
}

func (req *DescribeVpcsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeVpcsResponse, error) {
	resp := &DescribeVpcsResponse{}
	err := client.Request("vpc", "DescribeVpcs", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeVpcsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 符合条件的对象数。
	TotalCount int64 `json:"TotalCount"`
	// VPC对象。
	VpcSet []*Vpc `json:"VpcSet"`
}
