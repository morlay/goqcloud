package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询对端网关
// https://cloud.tencent.com/document/api/215/17516
type DescribeCustomerGatewaysRequest struct {
	// 对端网关ID，例如：cgw-2wqq41m9。每次请求的实例的上限为100。参数不支持同时指定CustomerGatewayIds和Filters。
	CustomerGatewayIds []*string `name:"CustomerGatewayIds,omitempty"`
	// 过滤条件，详见下表：实例过滤条件表。每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定CustomerGatewayIds和Filters。
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为20，最大值为100。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeCustomerGatewaysRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeCustomerGatewaysResponse, error) {
	resp := &DescribeCustomerGatewaysResponse{}
	err := client.Request("vpc", "DescribeCustomerGateways", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeCustomerGatewaysResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 对端网关对象列表
	CustomerGatewaySet []*CustomerGateway `json:"CustomerGatewaySet"`
	// 符合条件的实例数量。
	TotalCount int64 `json:"TotalCount"`
}
