package dc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建专用通道
// https://cloud.tencent.com/document/api/216/19821

type CreateDirectConnectTunnelRequest struct {
	// 专线带宽，单位：Mbps默认是物理专线带宽值
	Bandwidth *int64 `name:"Bandwidth,omitempty"`
	// BgpPeer，用户侧bgp信息，包括Asn和AuthKey
	BgpPeer *BgpPeer `name:"BgpPeer,omitempty"`
	// CustomerAddress，用户侧互联 IP
	CustomerAddress *string `name:"CustomerAddress,omitempty"`
	// 专线网关 ID，例如 dcg-d545ddf
	DirectConnectGatewayId *string `name:"DirectConnectGatewayId,omitempty"`
	// 专线 ID，例如：dc-kd7d06of
	DirectConnectId string `name:"DirectConnectId"`
	// 物理专线 owner，缺省为当前客户（物理专线 owner）共享专线时这里需要填写共享专线的开发商账号 ID
	DirectConnectOwnerAccount *string `name:"DirectConnectOwnerAccount,omitempty"`
	// 专用通道名称
	DirectConnectTunnelName string `name:"DirectConnectTunnelName"`
	// 网络地域
	NetworkRegion *string `name:"NetworkRegion,omitempty"`
	// 网络类型，分别为VPC、BMVPC，CCN，默认是VPCVPC：私有网络BMVPC：黑石网络CCN：云联网
	NetworkType *string `name:"NetworkType,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 静态路由，用户IDC的网段地址
	RouteFilterPrefixes []*RouteFilterPrefix `name:"RouteFilterPrefixes,omitempty"`
	// BGP ：BGP路由STATIC：静态默认为 BGP 路由
	RouteType *string `name:"RouteType,omitempty"`
	// TencentAddress，腾讯侧互联 IP
	TencentAddress *string `name:"TencentAddress,omitempty"`
	// vlan，范围：0 ~ 30000：不开启子接口默认值是非0
	Vlan *int64 `name:"Vlan,omitempty"`
	// 私有网络统一 ID 或者黑石网络统一 ID
	VpcId *string `name:"VpcId,omitempty"`
}

func (req *CreateDirectConnectTunnelRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateDirectConnectTunnelResponse, error) {
	resp := &CreateDirectConnectTunnelResponse{}
	err := client.Request("dc", "CreateDirectConnectTunnel", "2018-04-10").Do(req, resp)
	return resp, err
}

type CreateDirectConnectTunnelResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 专用通道ID
	DirectConnectTunnelIdSet []*string `json:"DirectConnectTunnelIdSet"`
}
