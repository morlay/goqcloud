package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 续费VPN网关
// https://cloud.tencent.com/document/api/215/17506

type RenewVpnGatewayRequest struct {
	// 预付费计费模式。
	InstanceChargePrepaid InstanceChargePrepaid `name:"InstanceChargePrepaid"`
	// 区域
	Region string `name:"Region"`
	// VPN网关实例ID。
	VpnGatewayId string `name:"VpnGatewayId"`
}

func (req *RenewVpnGatewayRequest) Invoke(client github_com_morlay_goqcloud.Client) (*RenewVpnGatewayResponse, error) {
	resp := &RenewVpnGatewayResponse{}
	err := client.Request("vpc", "RenewVpnGateway", "2017-03-12").Do(req, resp)
	return resp, err
}

type RenewVpnGatewayResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
