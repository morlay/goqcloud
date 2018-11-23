package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询专线网关
// https://cloud.tencent.com/document/api/215/30644

type DescribeDirectConnectGatewaysRequest struct {
	// 专线网关唯一ID，形如：dcg-9o233uri。
	DirectConnectGatewayIds []*string `name:"DirectConnectGatewayIds,omitempty"`
	// 过滤条件，参数不支持同时指定DirectConnectGatewayIds和Filters。direct-connect-gateway-id - String - 专线网关唯一ID，形如：dcg-9o233uri。direct-connect-gateway-name - String - 专线网关名称，默认模糊查询。direct-connect-gateway-ip - String - 专线网关IP。gateway-type - String - 网关类型，可选值：NORMAL（普通型）、NAT（NAT型）。network-type- String - 网络类型，可选值：VPC（私有网络类型）、CCN（云联网类型）。ccn-id - String - 专线网关所在私有网络ID。vpc-id - String - 专线网关所在云联网ID。
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeDirectConnectGatewaysRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDirectConnectGatewaysResponse, error) {
	resp := &DescribeDirectConnectGatewaysResponse{}
	err := client.Request("vpc", "DescribeDirectConnectGateways", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeDirectConnectGatewaysResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 专线网关对象数组。
	DirectConnectGatewaySet []*DirectConnectGateway `json:"DirectConnectGatewaySet"`
	// 符合条件的对象数。
	TotalCount int64 `json:"TotalCount"`
}
