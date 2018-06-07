package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除VPN通道
// https://cloud.tencent.com/document/api/215/17519

type DeleteVpnConnectionRequest struct {
	// 区域
	Region string `name:"Region"`
	// VPN通道实例ID。形如：vpnx-f49l6u0z。
	VpnConnectionId string `name:"VpnConnectionId"`
	// VPN网关实例ID。
	VpnGatewayId string `name:"VpnGatewayId"`
}

func (req *DeleteVpnConnectionRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteVpnConnectionResponse, error) {
	resp := &DeleteVpnConnectionResponse{}
	err := client.Request("vpc", "DeleteVpnConnection", "2017-03-12").Do(req, resp)
	return resp, err
}

type DeleteVpnConnectionResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
