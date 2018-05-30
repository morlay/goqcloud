package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询VPN通道列表
// https://cloud.tencent.com/document/api/215/17515
type DescribeVpnConnectionsRequest struct {
	// 过滤条件，详见下表：实例过滤条件表。每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定VpnConnectionIds和Filters。
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为20，最大值为100。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// VPN通道实例ID。形如：vpnx-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpnConnectionIds和Filters。
	VpnConnectionIds []*string `name:"VpnConnectionIds,omitempty"`
}

func (req *DescribeVpnConnectionsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeVpnConnectionsResponse, error) {
	resp := &DescribeVpnConnectionsResponse{}
	err := client.Request("vpc", "DescribeVpnConnections", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeVpnConnectionsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 符合条件的实例数量。
	TotalCount int64 `json:"TotalCount"`
	// VPN通道实例。
	VpnConnectionSet []*VpnConnection `json:"VpnConnectionSet"`
}
