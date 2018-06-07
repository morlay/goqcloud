package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 重置VPN通道
// https://cloud.tencent.com/document/api/215/17505

type ResetVpnConnectionRequest struct {
	// 区域
	Region string `name:"Region"`
	// VPN通道实例ID。形如：vpnx-f49l6u0z。
	VpnConnectionId string `name:"VpnConnectionId"`
	// VPN网关实例ID。
	VpnGatewayId string `name:"VpnGatewayId"`
}

func (req *ResetVpnConnectionRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ResetVpnConnectionResponse, error) {
	resp := &ResetVpnConnectionResponse{}
	err := client.Request("vpc", "ResetVpnConnection", "2017-03-12").Do(req, resp)
	return resp, err
}

type ResetVpnConnectionResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
