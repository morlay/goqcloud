package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 调整VPN网关带宽上限
// https://cloud.tencent.com/document/api/215/17504

type ResetVpnGatewayInternetMaxBandwidthRequest struct {
	// 公网带宽设置。可选带宽规格：5, 10, 20, 50, 100；单位：Mbps。
	InternetMaxBandwidthOut int64 `name:"InternetMaxBandwidthOut"`
	// 区域
	Region string `name:"Region"`
	// VPN网关实例ID。
	VpnGatewayId string `name:"VpnGatewayId"`
}

func (req *ResetVpnGatewayInternetMaxBandwidthRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ResetVpnGatewayInternetMaxBandwidthResponse, error) {
	resp := &ResetVpnGatewayInternetMaxBandwidthResponse{}
	err := client.Request("vpc", "ResetVpnGatewayInternetMaxBandwidth", "2017-03-12").Do(req, resp)
	return resp, err
}

type ResetVpnGatewayInternetMaxBandwidthResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
