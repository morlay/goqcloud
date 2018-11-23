package dc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询专用通道列表
// https://cloud.tencent.com/document/api/216/19819

type DescribeDirectConnectTunnelsRequest struct {
	// 专用通道 ID数组
	DirectConnectTunnelIds []*string `name:"DirectConnectTunnelIds,omitempty"`
	// 过滤条件:参数不支持同时指定DirectConnectTunnelIds和Filters。 direct-connect-tunnel-name, 专用通道名称。 direct-connect-tunnel-id, 专用通道实例ID，如dcx-abcdefgh。direct-connect-id, 物理专线实例ID，如，dc-abcdefgh。
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为20，最大值为100
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeDirectConnectTunnelsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDirectConnectTunnelsResponse, error) {
	resp := &DescribeDirectConnectTunnelsResponse{}
	err := client.Request("dc", "DescribeDirectConnectTunnels", "2018-04-10").Do(req, resp)
	return resp, err
}

type DescribeDirectConnectTunnelsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 专用通道列表
	DirectConnectTunnelSet []*DirectConnectTunnel `json:"DirectConnectTunnelSet"`
	// 符合专用通道数量。
	TotalCount int64 `json:"TotalCount"`
}
