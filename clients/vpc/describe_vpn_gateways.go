package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询VPN网关
// https://cloud.tencent.com/document/api/215/17514

type DescribeVpnGatewaysRequest struct {
	// 过滤条件，参数不支持同时指定VpnGatewayIds和Filters。vpc-id - String - （过滤条件）VPC实例ID形如：vpc-f49l6u0z。vpn-gateway-id - String - （过滤条件）VPN实例ID形如：vpngw-5aluhh9t。vpn-gateway-name - String - （过滤条件）VPN实例名称。type - String - （过滤条件）VPN网关类型：'IPSEC', 'SSL'。public-ip-address- String - （过滤条件）公网IP。renew-flag - String - （过滤条件）网关续费类型，手动续费：'NOTIFY_AND_MANUAL_RENEW'、自动续费：'NOTIFY_AND_AUTO_RENEW'。zone - String - （过滤条件）VPN所在可用区，形如：ap-guangzhou-2。
	Filters []*FilterObject `name:"Filters,omitempty"`
	// 请求对象个数
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// VPN网关实例ID。形如：vpngw-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpnGatewayIds和Filters。
	VpnGatewayIds []*string `name:"VpnGatewayIds,omitempty"`
}

func (req *DescribeVpnGatewaysRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeVpnGatewaysResponse, error) {
	resp := &DescribeVpnGatewaysResponse{}
	err := client.Request("vpc", "DescribeVpnGateways", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeVpnGatewaysResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 符合条件的实例数量。
	TotalCount int64 `json:"TotalCount"`
	// VPN网关实例详细信息列表。
	VpnGatewaySet []*VpnGateway `json:"VpnGatewaySet"`
}
