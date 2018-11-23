package dc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改专用通道属性
// https://cloud.tencent.com/document/api/216/19818

type ModifyDirectConnectTunnelAttributeRequest struct {
	// 专用通道带宽值，单位为M。
	Bandwidth *int64 `name:"Bandwidth,omitempty"`
	// 用户侧BGP，包括Asn，AuthKey
	BgpPeer *BgpPeer `name:"BgpPeer,omitempty"`
	// 用户侧互联IP
	CustomerAddress *string `name:"CustomerAddress,omitempty"`
	// 专用通道ID
	DirectConnectTunnelId string `name:"DirectConnectTunnelId"`
	// 专用通道名称
	DirectConnectTunnelName *string `name:"DirectConnectTunnelName,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 用户侧网段地址
	RouteFilterPrefixes []*RouteFilterPrefix `name:"RouteFilterPrefixes,omitempty"`
	// 腾讯侧互联IP
	TencentAddress *string `name:"TencentAddress,omitempty"`
}

func (req *ModifyDirectConnectTunnelAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyDirectConnectTunnelAttributeResponse, error) {
	resp := &ModifyDirectConnectTunnelAttributeResponse{}
	err := client.Request("dc", "ModifyDirectConnectTunnelAttribute", "2018-04-10").Do(req, resp)
	return resp, err
}

type ModifyDirectConnectTunnelAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
