package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除VPN网关
// https://cloud.tencent.com/document/api/215/17518
type DeleteVpnGatewayRequest struct {
	// 区域
	Region string `name:"Region"`
	// VPN网关实例ID。
	VpnGatewayId string `name:"VpnGatewayId"`
}

func (req *DeleteVpnGatewayRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteVpnGatewayResponse, error) {
	resp := &DeleteVpnGatewayResponse{}
	err := client.Request("vpc", "DeleteVpnGateway", "2017-03-12").Do(req, resp)
	return resp, err
}

type DeleteVpnGatewayResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
